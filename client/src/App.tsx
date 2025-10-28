import { useEffect, useState } from "react";
import {
  getPremades,
  receiveBattlers,
  sendBattleChoice,
  sendChoice,
  sendCustom,
} from "./api/api";
import Battler from "./components/Battler";
import ListChoice from "./components/ListChoice";
import Header from "./components/Header";
import { Button } from "./components/ui/button";
import { Skeleton } from "./components/ui/skeleton";
import { ThemeProvider } from "./components/theme-provider";
import { Alert, AlertDescription, AlertTitle } from "./components/ui/alert";

// Battler type for round battlers
export type Battlers = {
  Name: string;
  Score: number;
};

export default function App() {
  const [premadeLists, setPremadeLists] = useState<string[]>([]);
  const [currentBattlers, setCurrentBattlers] = useState<Battlers[]>([]);
  const [finalRanking, setFinalRanking] = useState<string[]>([]);
  const [listLoading, setListLoading] = useState<boolean>(true);
  const [gameStart, setGameStart] = useState<boolean>(false);
  const [itemsLeft, setItemsLeft] = useState<number>(0);
  const [roundRobin, setRoundRobin] = useState<boolean>(false);
  const [customList, setCustomList] = useState<string>("");
  const [validationError, setValidationError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPremades = async () => {
      const data = await getPremades();
      if (data) setPremadeLists(data.premades);
      setListLoading(false);
    };
    fetchPremades();
  }, []);
  // Take array of premade lists and map them out to buttons for user to select list choice
  const premadeOptions = listLoading
    ? Array.from({ length: 2 }).map((_, i) => (
        <Skeleton
          key={i}
          className="m-2.5 h-24 w-fit min-w-[15rem] bg-gray-200 px-10"
        />
      ))
    : premadeLists.map((item: string) => (
        <ListChoice
          key={item}
          listSelection={async () => {
            await sendChoice(item);
            receiveBattlers(
              setCurrentBattlers,
              setFinalRanking,
              setItemsLeft,
              setRoundRobin,
            );
            setGameStart(true);
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
        receiveBattlers(
          setCurrentBattlers,
          setFinalRanking,
          setItemsLeft,
          setRoundRobin,
        );
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
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <main className="container mx-auto min-h-screen p-4">
        <Header />
        {!gameStart && (
          <div className="mt-10 flex flex-col items-center gap-4">
            <h2 className="text-my-color text-center text-xl font-semibold md:text-3xl">
              Choose a premade list to start battling
            </h2>
            <ul className="flex list-none flex-col sm:items-center md:flex-row">
              {premadeOptions}
            </ul>
          </div>
        )}

        {currentBattlers.length > 0 && (
          <div className="mt-6 flex flex-col items-center gap-4">
            {!roundRobin ? (
              <h2 className="text-my-color text-center text-3xl">
                Choose Which Item You Prefer
              </h2>
            ) : (
              <>
                {" "}
                <h2 className="text-my-color text-center text-3xl">
                  Round Robin!
                </h2>
                <h3 className="text-my-color text-center text-2xl">
                  Each remaining item will battle to determine the final ranking
                </h3>
              </>
            )}
            <h2 className="text-my-color text-center text-2xl">
              Items Left: {itemsLeft.toString()}
            </h2>
            <div className="mt-10 flex flex-col place-items-center gap-10 px-5 xl:flex-row">
              {battleOptions}
            </div>
            {!roundRobin && (
              <Button
                onClick={() =>
                  receiveBattlers(
                    setCurrentBattlers,
                    setFinalRanking,
                    setItemsLeft,
                    setRoundRobin,
                  )
                }
                className="choice-button mt-2 cursor-pointer p-8 text-center text-lg shadow-sky-200 transition-colors duration-200 md:text-xl lg:mt-5"
                variant={"ghost"}
              >
                Can't choose? Skip round
              </Button>
            )}
          </div>
        )}

        <div className="flex flex-col">
          <label htmlFor="custom list">
            Please input items for your custom list, separated by a new line
          </label>
          <textarea
            name="custom list"
            id="custom"
            onChange={(e) => setCustomList(e.target.value)}
            className="w-1/2 border-2 border-gray-100"
          ></textarea>
          <Button
            className="w-1/2"
            onClick={async () => {
              const success = await sendCustom(customList, setValidationError);
              if (success) {
                receiveBattlers(
                  setCurrentBattlers,
                  setFinalRanking,
                  setItemsLeft,
                  setRoundRobin,
                );
              }
            }}
          >
            Play
          </Button>
          {validationError && (
            <Alert variant={"destructive"} className="w-1/2">
              <AlertTitle>Oops!</AlertTitle>
              <AlertDescription>{validationError}</AlertDescription>
            </Alert>
          )}
        </div>

        {finalRanking.length > 0 && (
          <div className="mt-6 flex flex-col items-center">
            <h2 className="text-my-color text-4xl font-bold md:text-5xl">
              Final Results
            </h2>
            <div className="text-my-color mt-3 flex flex-col gap-3 text-center md:mt-6">
              {finalRanking.length > 0 && rankingList}
            </div>
          </div>
        )}
      </main>
    </ThemeProvider>
  );
}
