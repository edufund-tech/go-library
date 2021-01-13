package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// User domain model
type User struct {
	UID           string `bson:"uid,omitempty" json:"uid,omitempty"`
	Email         string `bson:"email,omitempty" json:"email,omitempty"`
	EmailVerified bool   `bson:"email_verified,omitempty" json:"email_verified,omitempty"`
	PhoneNumber   string `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	DisplayName   string `bson:"display_name,omitempty" json:"display_name,omitempty"`
	PhotoURL      string `bson:"photo_url,omitempty" json:"photo_url,omitempty"`
	Disabled      bool   `bson:"disabled,omitempty" json:"disabled,omitempty"`
	Password      string `json:"password,omitempty"`
}

// Client ...
type Client struct {
	*auth.Client
}

func Connect(ctx context.Context, serviceAccount string) (client *Client, err error) {
	opt := option.WithCredentialsFile(serviceAccount)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error connecting firebase: %v\n", err)
	}

	auth, err := app.Auth(ctx)
	client.Client = auth
	if err != nil {
		log.Fatalf("error client spawn: %v\n", err)
	}
	return client, err
}

func (client *Client) Create(ctx context.Context, user User) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		DisplayName(user.DisplayName).
		Password(user.Password).
		PhotoURL(user.PhotoURL).
		Disabled(false)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	return u, err
}

func (client *Client) Update(ctx context.Context, user User) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		DisplayName(user.DisplayName).
		Password(user.Password).
		PhotoURL(user.PhotoURL).
		Disabled(false)
	u, err := client.UpdateUser(ctx, user.UID, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	return u, err
}

func (client *Client) Delete(ctx context.Context, user User) error {
	err := client.DeleteUser(ctx, user.UID)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted user: %v\n", user)
	return err
}

func (client *Client) VerifyToken(ctx context.Context, idToken string) (verified bool, err error) {
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return true, err
}
