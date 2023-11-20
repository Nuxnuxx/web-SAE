import express from "express"
import morgan from "morgan"
import recipeRouter from "./api/v1/router/recipe.router.js"

const app = express()

app.use(morgan("dev"))

app.use("/api/v1/recipe", recipeRouter)

export default app
