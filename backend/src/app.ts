import express from "express"
import recipeRouter from "./api/v1/router/recipe.router.js"

const app = express()

app.use("/api/v1/recipe", recipeRouter)

export default app
