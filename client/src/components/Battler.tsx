import { Card, CardContent } from "@/components/ui/card";

interface BattlerProps {
  battlerName: string;
  winner: () => Promise<void>;
}

export default function Battler({
  battlerName = "Item",
  winner,
}: BattlerProps) {
  return (
    <Card
      className="flex min-h-[15rem] w-full min-w-[16rem] cursor-pointer items-center justify-center bg-sky-200 text-center shadow-sky-200 transition-colors duration-200 hover:bg-sky-400 sm:w-[32rem]"
      onClick={winner}
    >
      <CardContent className="text-2xl text-sky-950 sm:text-4xl">
        {battlerName}
      </CardContent>
    </Card>
  );
}
