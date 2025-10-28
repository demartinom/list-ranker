import { useState } from "react";
import { Button } from "./ui/button";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { sendCustom, receiveBattlers } from "@/api/api";
import { Battlers } from "@/App";

interface CustomListProps {
  setCurrentBattlers: React.Dispatch<React.SetStateAction<Battlers[]>>;
  setFinalRanking: React.Dispatch<React.SetStateAction<string[]>>;
  setItemsLeft: React.Dispatch<React.SetStateAction<number>>;
  setRoundRobin: React.Dispatch<React.SetStateAction<boolean>>;
  setGameStart: React.Dispatch<React.SetStateAction<boolean>>;
}

export default function CustomList({
  setCurrentBattlers,
  setFinalRanking,
  setItemsLeft,
  setRoundRobin,
  setGameStart,
}: CustomListProps) {
  const [customList, setCustomList] = useState<string>("");
  const [validationError, setValidationError] = useState<string | null>(null);

  return (
    <div className="mt-10 flex flex-col items-center">
      <label
        htmlFor="custom list"
        className="text-my-color text-center text-2xl font-semibold"
      >
        Want to use your own custom list?
        <br /> Input list items below using a new line to separate items.
      </label>
      <textarea
        name="custom list"
        id="custom"
        onChange={(e) => setCustomList(e.target.value)}
        className="w-1/2 border-2 border-gray-100"
      ></textarea>
      <Button
        variant={"ghost"}
        className="choice-button text-my-color w-1/2 py-6 text-lg"
        onClick={async () => {
          const success = await sendCustom(customList, setValidationError);
          if (success) {
            receiveBattlers(
              setCurrentBattlers,
              setFinalRanking,
              setItemsLeft,
              setRoundRobin,
            );
            setGameStart(true);
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
  );
}
