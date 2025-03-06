import { useEffect, useState } from "react";
import {
  getPremades,
  receiveBattlers,
  sendBattleChoice,
  sendChoice,
} from "./api/api";

export type Battlers = {
  Name: string;
  Score: number;
};
export default function App() {
  const [premadeLists, setPremadeLists] = useState<string[]>([]);
  const [currentBattlers, setCurrentBattlers] = useState<Battlers[]>([]);

  useEffect(() => {
    const fetchPremades = async () => {
      const data = await getPremades();
      if (data) setPremadeLists(data.premades);
    };
    fetchPremades();
  }, []);

  // Take array of premade lists and map them out to buttons for user to select list choice
  const premadeOptions = premadeLists.map((item: string) => (
    <button
      key={item}
      onClick={async () => {
        await sendChoice(item);
        receiveBattlers(setCurrentBattlers);
      }}
    >
      {item[0].toUpperCase() + item.substring(1)}
    </button>
  ));

  // Take list of round battlers and display as choices
  const battleOptions = currentBattlers.map(
    (battler: Battlers, index: number) => (
      <button
        key={index}
        onClick={async () => {
          await sendBattleChoice(battler.Name);
          receiveBattlers(setCurrentBattlers);
        }}
      >
        {battler.Name}
      </button>
    )
  );

  return (
    <div>
      {premadeLists.length > 0 && premadeOptions}
      {currentBattlers.length > 0 && battleOptions}
    </div>
  );
}
