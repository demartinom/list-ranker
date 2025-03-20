import { useEffect, useState } from "react";
import {
  getPremades,
  receiveBattlers,
  sendBattleChoice,
  sendChoice,
} from "./api/api";
import Battler from "./components/Battler";

export type Battlers = {
  Name: string;
  Score: number;
};
export default function App() {
  const [premadeLists, setPremadeLists] = useState<string[]>([]);
  const [currentBattlers, setCurrentBattlers] = useState<Battlers[]>([]);
  const [finalRanking, setFinalRanking] = useState<string[]>([]);

  useEffect(() => {
    const fetchPremades = async () => {
      const data = await getPremades();
      if (data) setPremadeLists(data.premades);
    };
    fetchPremades();
  }, []);

  // Take array of premade lists and map them out to buttons for user to select list choice
  const premadeOptions = premadeLists.map((item: string) => (
    <ListChoice
      key={item}
      listSelection={async () => {
        await sendChoice(item);
        receiveBattlers(setCurrentBattlers, setFinalRanking);
      }}
      listName={item[0].toUpperCase() + item.substring(1)}
    />
  ));

  // Take list of round battlers and display as choices
  const battleOptions = currentBattlers.map(
    (battler: Battlers, ) => (
      <Battler 
      key={battler.Name}
      name={battler.Name}
      winner={async () => {
        await sendBattleChoice(battler.Name);
        receiveBattlers(setCurrentBattlers, setFinalRanking);
      }}
    />
    )
  );

  const rankingList = finalRanking.map((item: string, index: number) => (
    <p key={index}>
      {index + 1}: {item}
    </p>
  ));

  return (
    <div>
      {premadeLists.length > 0 && premadeOptions}
      {currentBattlers.length > 0 && battleOptions}
      {finalRanking.length > 0 && rankingList}
    </div>
  );
}
