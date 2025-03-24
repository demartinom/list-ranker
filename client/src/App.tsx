import { useEffect, useState } from "react";
import {
  getPremades,
  receiveBattlers,
  sendBattleChoice,
  sendChoice,
} from "./api/api";
import Battler from "./components/Battler";
import ListChoice from "./components/ListChoice";
import { Separator } from "./components/ui/separator";

// Battler type for round battlers
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

  // Map final rankings to list with rank number
  const rankingList = finalRanking.map((item: string, index: number) => (
    <p key={index} className="text-3xl">
      {index + 1}: {item}
    </p>
  ));

  return (
    <main className="m-auto min-h-screen w-5/6 max-w-4xl px-4">
      <div className="mt-8 flex flex-col items-center">
        <h1 className="text-7xl font-extrabold text-gray-700">List Ranker</h1>
        <Separator className="mt-5" style={{ height: "3px" }} />
      </div>
      {premadeLists.length > 0 && (
        <div className="mt-10 flex flex-col items-center gap-2 p-4 sm:gap-4">
          <h2 className="text-3xl font-semibold text-gray-700">
            Choose a premade list to start battling
          </h2>
          <ul className="list-none">{premadeOptions}</ul>
        </div>
      )}
      {currentBattlers.length > 0 && (
        <div className="mt-6 flex flex-col items-center gap-2 p-4 sm:gap-4">
          <h2 className="text-3xl text-gray-700">
            Choose Which Item You Prefer
          </h2>
          <div className="mt-10 flex justify-center gap-10">
            {battleOptions}
          </div>
        </div>
      )}
      {finalRanking.length > 0 && (
        <div className="mt-12 flex flex-col items-center pt-10">
          <h2 className="text-5xl font-bold text-gray-900">Final Results</h2>
          <div className="mt-6 flex flex-col gap-3 text-center">
            {finalRanking.length > 0 && rankingList}
          </div>
        </div>
      )}
    </main>
  );
}
