export const sendLogin = async (email: string, password: string) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/api/v1/auth/login`,
			{
				method: "POST",
				body: JSON.stringify({ email, password }),
				headers: {
					"Content-type": "application/json",
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { message } = await result.json();
			return message;
		}
	} catch (err) {
		throw err;
	}
};

export const sendRegister = async (
	firstName: string,
	lastName: string,
	email: string,
	password: string
) => {
	try {
		const name = firstName + " " + lastName;
		const result = await fetch(
			"http://localhost:5000/api/v1/auth/register",
			{
				method: "POST",
				body: JSON.stringify({ name, email, password }),
				headers: {
					"Content-type": "application/json",
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			throw new Error("What is this shit");
		}
	} catch (err) {
		throw new Error("Doesnt work this shitty api");
	}
};
