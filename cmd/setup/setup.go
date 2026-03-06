package setup

import "profitti/internal/app/transport/http/handlers/users"

type Setup struct {
	UserHandler users.UserHandler
}
