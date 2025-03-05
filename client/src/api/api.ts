import axios from "axios";
import { Battlers } from "../App";
const apiPath: string = "http://localhost:8080/api";

export const getPremades = async () => {
  try {
    const response = await axios.get(`${apiPath}/premades`);
    return response.data;
  } catch (error) {
    console.error("API error:, ", error);
    return null;
  }
};

export const sendChoice = async (choice: string) => {
  const message = JSON.stringify({ selection: choice });
  await axios.post(`${apiPath}/listchoice`, message);
};

export const receiveBattlers = async (
  listSetter: React.Dispatch<React.SetStateAction<Battlers[]>>
) => {
  try {
    const response = await axios.post(`${apiPath}/battlers`);
    listSetter(response.data.battlers);
  } catch (error) {
    console.log(error);
  }
};

export const sendBattleChoice = async (choice: string) => {
  const message = JSON.stringify({ selection: choice });
  await axios.post(`${apiPath}/battlerChoice`, message);
};
