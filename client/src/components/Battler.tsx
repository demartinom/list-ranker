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
      className="choice-button light:shadow-sky-200 flex w-full min-w-[16rem] cursor-pointer items-center justify-center text-center transition-colors duration-200 sm:w-[32rem] md:min-h-[15rem]"
      onClick={winner}
    >
      <CardContent className="text-2xl sm:text-4xl">{battlerName}</CardContent>
    </Card>
  );
}
