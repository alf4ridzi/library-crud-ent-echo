package tokenutil

import "github.com/alf4ridzi/library-crud-ent-echo/internal/config"

var JwtAuthSecret = config.AppConfig.JwtAuthSecret
var JwtRefreshSecret = config.AppConfig.JwtRefreshSecret
