import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "./ui/skeleton";

interface BattlerProps {
  battlerName: string;
  winner: () => Promise<void>;
}

export default function Battler({
  battlerName = "Item",
  winner,
}: BattlerProps) {
  return (
    <>
      {!battlerName ? (
        <Skeleton className="h-60 w-full justify-center bg-gray-200 sm:w-2xl" />
      ) : (
        <Card
          className="h-60 w-full cursor-pointer justify-center bg-sky-200 text-center shadow-sky-200 transition-colors duration-200 hover:bg-sky-400 sm:w-2xl"
          onClick={winner}
        >
          <CardContent className="text-4xl text-sky-950">
            {battlerName}
          </CardContent>
        </Card>
      )}
    </>
  );
}
