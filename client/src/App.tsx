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
    <p key={index} className="text-2xl md:text-3xl">
      {index + 1}: {item}
    </p>
  ));

  return (
    <main className="container mx-auto min-h-screen p-4">
      <div className="mt-8 flex flex-col items-center text-center">
        <h1 className="text-5xl font-extrabold text-gray-700 md:text-7xl">
          List Ranker
        </h1>
        <Separator className="mt-5" style={{ height: "3px" }} />
      </div>

      {premadeLists.length > 0 && (
        <div className="mt-10 flex flex-col items-center gap-4">
          <h2 className="text-3xl font-semibold text-gray-700">
            Choose a premade list to start battling
          </h2>
          <ul className="list-none">{premadeOptions}</ul>
        </div>
      )}

      {currentBattlers.length > 0 && (
        <div className="mt-6 flex flex-col items-center gap-4">
          <h2 className="text-center text-3xl text-gray-700">
            Choose Which Item You Prefer
          </h2>
          <div className="mt-10 flex flex-col place-items-center gap-10 px-5 xl:flex-row">
            {battleOptions}
          </div>
        </div>
      )}

      {finalRanking.length > 0 && (
        <div className="mt-6 flex flex-col items-center">
          <h2 className="text-4xl font-bold text-gray-900 md:text-5xl">
            Final Results
          </h2>
          <div className="mt-3 flex flex-col gap-3 text-center md:mt-6">
            {finalRanking.length > 0 && rankingList}
          </div>
        </div>
      )}
    </main>
  );
}
