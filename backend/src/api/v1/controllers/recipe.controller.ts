import { Request, Response } from "express"

export const getProduct = async (req: Request, res: Response) => {
	res.send("Hello from controller")
}
