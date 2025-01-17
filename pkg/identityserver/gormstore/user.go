// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// User model.
type User struct {
	Model
	SoftDelete

	Account Account `gorm:"polymorphic:Account;polymorphic_value:user"`

	// BEGIN common fields
	Name        string      `gorm:"type:VARCHAR"`
	Description string      `gorm:"type:TEXT"`
	Attributes  []Attribute `gorm:"polymorphic:Entity;polymorphic_value:user"`
	APIKeys     []APIKey    `gorm:"polymorphic:Entity;polymorphic_value:user"`
	// END common fields

	Sessions []*UserSession

	PrimaryEmailAddress            string     `gorm:"type:VARCHAR;not null;unique_index"`
	PrimaryEmailAddressValidatedAt *time.Time // should be cleared when email changes

	Password              string    `gorm:"type:VARCHAR;not null"` // this is the hash
	PasswordUpdatedAt     time.Time `gorm:"not null"`
	RequirePasswordUpdate bool      `gorm:"not null"`

	State            int    `gorm:"not null"`
	StateDescription string `gorm:"type:VARCHAR"`

	Admin bool `gorm:"not null"`

	TemporaryPassword          string `gorm:"type:VARCHAR"`
	TemporaryPasswordCreatedAt *time.Time
	TemporaryPasswordExpiresAt *time.Time

	ProfilePicture   *Picture
	ProfilePictureID *string `gorm:"type:UUID;index:user_profile_picture_index"`
}

func init() {
	registerModel(&User{})
}

// functions to set fields from the user model into the user proto.
var userPBSetters = map[string]func(*ttnpb.User, *User){
	nameField:                func(pb *ttnpb.User, usr *User) { pb.Name = usr.Name },
	descriptionField:         func(pb *ttnpb.User, usr *User) { pb.Description = usr.Description },
	attributesField:          func(pb *ttnpb.User, usr *User) { pb.Attributes = attributes(usr.Attributes).toMap() },
	primaryEmailAddressField: func(pb *ttnpb.User, usr *User) { pb.PrimaryEmailAddress = usr.PrimaryEmailAddress },
	primaryEmailAddressValidatedAtField: func(pb *ttnpb.User, usr *User) {
		pb.PrimaryEmailAddressValidatedAt = ttnpb.ProtoTime(cleanTimePtr(usr.PrimaryEmailAddressValidatedAt))
	},
	passwordField: func(pb *ttnpb.User, usr *User) { pb.Password = usr.Password },
	passwordUpdatedAtField: func(pb *ttnpb.User, usr *User) {
		pb.PasswordUpdatedAt = ttnpb.ProtoTimePtr(cleanTime(usr.PasswordUpdatedAt))
	},
	requirePasswordUpdateField: func(pb *ttnpb.User, usr *User) { pb.RequirePasswordUpdate = usr.RequirePasswordUpdate },
	stateField:                 func(pb *ttnpb.User, usr *User) { pb.State = ttnpb.State(usr.State) },
	stateDescriptionField:      func(pb *ttnpb.User, usr *User) { pb.StateDescription = usr.StateDescription },
	adminField:                 func(pb *ttnpb.User, usr *User) { pb.Admin = usr.Admin },
	temporaryPasswordField: func(pb *ttnpb.User, usr *User) {
		pb.TemporaryPassword = usr.TemporaryPassword
	},
	temporaryPasswordCreatedAtField: func(pb *ttnpb.User, usr *User) {
		pb.TemporaryPasswordCreatedAt = ttnpb.ProtoTime(cleanTimePtr(usr.TemporaryPasswordCreatedAt))
	},
	temporaryPasswordExpiresAtField: func(pb *ttnpb.User, usr *User) {
		pb.TemporaryPasswordExpiresAt = ttnpb.ProtoTime(cleanTimePtr(usr.TemporaryPasswordExpiresAt))
	},
	profilePictureField: func(pb *ttnpb.User, usr *User) {
		if usr.ProfilePicture == nil {
			pb.ProfilePicture = nil
		} else {
			pb.ProfilePicture = usr.ProfilePicture.toPB()
		}
	},
}

// functions to set fields from the user proto into the user model.
var userModelSetters = map[string]func(*User, *ttnpb.User){
	nameField:        func(usr *User, pb *ttnpb.User) { usr.Name = pb.Name },
	descriptionField: func(usr *User, pb *ttnpb.User) { usr.Description = pb.Description },
	attributesField: func(usr *User, pb *ttnpb.User) {
		usr.Attributes = attributes(usr.Attributes).updateFromMap(pb.Attributes)
	},
	primaryEmailAddressField: func(usr *User, pb *ttnpb.User) { usr.PrimaryEmailAddress = pb.PrimaryEmailAddress },
	primaryEmailAddressValidatedAtField: func(usr *User, pb *ttnpb.User) {
		usr.PrimaryEmailAddressValidatedAt = ttnpb.StdTime(pb.PrimaryEmailAddressValidatedAt)
	},
	passwordField: func(usr *User, pb *ttnpb.User) { usr.Password = pb.Password },
	passwordUpdatedAtField: func(usr *User, pb *ttnpb.User) {
		if passwordUpdatedAt := ttnpb.StdTime(pb.PasswordUpdatedAt); passwordUpdatedAt != nil {
			usr.PasswordUpdatedAt = cleanTime(*passwordUpdatedAt)
		} else {
			usr.PasswordUpdatedAt = cleanTime(time.Now())
		}
	},
	requirePasswordUpdateField: func(usr *User, pb *ttnpb.User) { usr.RequirePasswordUpdate = pb.RequirePasswordUpdate },
	stateField:                 func(usr *User, pb *ttnpb.User) { usr.State = int(pb.State) },
	stateDescriptionField:      func(usr *User, pb *ttnpb.User) { usr.StateDescription = pb.StateDescription },
	adminField:                 func(usr *User, pb *ttnpb.User) { usr.Admin = pb.Admin },
	temporaryPasswordField: func(usr *User, pb *ttnpb.User) {
		usr.TemporaryPassword = pb.TemporaryPassword
	},
	temporaryPasswordCreatedAtField: func(usr *User, pb *ttnpb.User) {
		usr.TemporaryPasswordCreatedAt = cleanTimePtr(ttnpb.StdTime(pb.TemporaryPasswordCreatedAt))
	},
	temporaryPasswordExpiresAtField: func(usr *User, pb *ttnpb.User) {
		usr.TemporaryPasswordExpiresAt = cleanTimePtr(ttnpb.StdTime(pb.TemporaryPasswordExpiresAt))
	},
	profilePictureField: func(usr *User, pb *ttnpb.User) {
		usr.ProfilePictureID, usr.ProfilePicture = nil, nil
		if pb.ProfilePicture != nil {
			usr.ProfilePicture = &Picture{}
			usr.ProfilePicture.fromPB(pb.ProfilePicture)
		}
	},
}

// fieldMask to use if a nil or empty fieldmask is passed.
var defaultUserFieldMask store.FieldMask

func init() {
	paths := make([]string, 0, len(userPBSetters))
	for _, path := range ttnpb.UserFieldPathsNested {
		if _, ok := userPBSetters[path]; ok {
			paths = append(paths, path)
		}
	}
	defaultUserFieldMask = paths
}

// fieldmask path to column name in users table.
var userColumnNames = map[string][]string{
	attributesField:                     {},
	contactInfoField:                    {},
	nameField:                           {nameField},
	descriptionField:                    {descriptionField},
	primaryEmailAddressField:            {primaryEmailAddressField},
	primaryEmailAddressValidatedAtField: {primaryEmailAddressValidatedAtField},
	passwordField:                       {passwordField},
	passwordUpdatedAtField:              {passwordUpdatedAtField},
	requirePasswordUpdateField:          {requirePasswordUpdateField},
	stateField:                          {stateField},
	stateDescriptionField:               {stateDescriptionField},
	adminField:                          {adminField},
	temporaryPasswordField:              {temporaryPasswordField},
	temporaryPasswordCreatedAtField:     {temporaryPasswordCreatedAtField},
	temporaryPasswordExpiresAtField:     {temporaryPasswordExpiresAtField},
}

func (usr User) toPB(pb *ttnpb.User, fieldMask store.FieldMask) {
	pb.Ids = &ttnpb.UserIdentifiers{UserId: usr.Account.UID}
	pb.CreatedAt = ttnpb.ProtoTimePtr(cleanTime(usr.CreatedAt))
	pb.UpdatedAt = ttnpb.ProtoTimePtr(cleanTime(usr.UpdatedAt))
	pb.DeletedAt = ttnpb.ProtoTime(cleanTimePtr(usr.DeletedAt))
	if len(fieldMask) == 0 {
		fieldMask = defaultUserFieldMask
	}
	for _, path := range fieldMask {
		if setter, ok := userPBSetters[path]; ok {
			setter(pb, &usr)
		}
	}
}

func (usr *User) fromPB(pb *ttnpb.User, fieldMask store.FieldMask) (columns []string) {
	if len(fieldMask) == 0 {
		fieldMask = defaultUserFieldMask
	}
	for _, path := range fieldMask {
		if setter, ok := userModelSetters[path]; ok {
			setter(usr, pb)
			if columnNames, ok := userColumnNames[path]; ok {
				columns = append(columns, columnNames...)
			}
			continue
		}
	}
	return columns
}

type userWithUID struct {
	UID  string
	User `gorm:"embedded"`
}

func (userWithUID) TableName() string { return "users" }

func (u userWithUID) toPB(pb *ttnpb.User, fieldMask store.FieldMask) {
	u.User.Account.UID = u.UID
	u.User.toPB(pb, fieldMask)
}
