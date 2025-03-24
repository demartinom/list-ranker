import { useEffect, useState } from "react";
import {
  getPremades,
  receiveBattlers,
  sendBattleChoice,
  sendChoice,
} from "./api/api";
import Battler from "./components/Battler";
import ListChoice from "./components/ListChoice";

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
        setPremadeLists([]);
      }}
      listName={item[0].toUpperCase() + item.substring(1)}
    />
  ));

  // Take list of round battlers and display as choices
  const battleOptions = currentBattlers.map((battler: Battlers) => (
    <Battler
      key={battler.Name}
      battlerName={battler.Name}
      winner={async () => {
        await sendBattleChoice(battler.Name);
        receiveBattlers(setCurrentBattlers, setFinalRanking);
      }}
    />
  ));

  const rankingList = finalRanking.map((item: string, index: number) => (
    <p key={index} className="text-3xl">
      {index + 1}: {item}
    </p>
  ));

  return (
    <main className="m-auto min-h-screen w-5/6">
      {premadeLists.length > 0 && (
        <div className="flex flex-col items-center gap-2 p-4 sm:gap-4">
          <h2 className="text-3xl">Choose a premade list to start battling</h2>
          <ul className="list-none">{premadeOptions}</ul>
        </div>
      )}
      {currentBattlers.length > 0 && (
        <div className="flex flex-col items-center gap-2 p-4 sm:gap-4">
          <h2 className="text-3xl">Choose Which Item You Prefer</h2>
          <div className="flex justify-center gap-10">{battleOptions}</div>
        </div>
      )}
      {finalRanking.length > 0 && (
        <div className="mt-16 flex flex-col items-center">
          <h2 className="text-5xl">Final Results</h2>
          <div className="mt-8 flex flex-col gap-2">
            {finalRanking.length > 0 && rankingList}
          </div>
        </div>
      )}
    </main>
  );
}
