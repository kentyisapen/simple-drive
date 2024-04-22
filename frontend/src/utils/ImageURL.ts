export const getImageURL = (uuid: string) => {
	if (uuid === "00000000-0000-0000-0000-000000000000") return "/no_image.png";
	return `http://localhost:9000/images/${uuid}`;
};
