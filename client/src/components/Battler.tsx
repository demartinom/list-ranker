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
        <Skeleton className="min-h-[10rem] w-full min-w-[16rem] bg-gray-200 sm:w-[32rem]" />
      ) : (
        <Card
          className="flex min-h-[15rem] w-full min-w-[16rem] cursor-pointer items-center justify-center bg-sky-200 text-center shadow-sky-200 transition-colors duration-200 hover:bg-sky-400 sm:w-[32rem]"
          onClick={winner}
        >
          <CardContent className="text-2xl text-sky-950 sm:text-4xl">
            {battlerName}
          </CardContent>
        </Card>
      )}
    </>
  );
}
