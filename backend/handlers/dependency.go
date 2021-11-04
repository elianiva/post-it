package handlers

import "post-it-backend/prisma/db"

type Dependency struct {
	DB *db.PrismaClient
}
