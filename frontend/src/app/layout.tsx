import { store } from "@/store/store";
import { Container } from "@mui/material";
import type { Metadata } from "next";

export const metadata: Metadata = {
	title: "Create Next App",
	description: "Generated by create next app",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en">
			<body>
				<Container maxWidth={false}>
					<div>{children}</div>
				</Container>
			</body>
		</html>
	);
}
