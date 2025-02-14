import axios from "axios";

export const getPremades = async () => {
  try {
    const response = await axios.get("http://localhost:8080/api/premades");
    return response.data;
  } catch (error) {
    console.error("API error:, ", error);
    return null;
  }
};

export const sendChoice = async (choice: string) => {
  const message = JSON.stringify({ selection: choice });
  await axios.post("http://localhost:8080/api/listchoice", message);
};
