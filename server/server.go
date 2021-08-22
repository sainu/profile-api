package server

import "github.com/sainu/profile-api/config"

// Init initialize server
func Init() error {
	c := config.GetConfig()
	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Logger.Info(r.Start(":" + c.GetString("server.port")))
	return nil
}
