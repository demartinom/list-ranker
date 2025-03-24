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
      className="h-60 w-2xl cursor-pointer justify-center bg-sky-200 text-center shadow-sky-200 transition-colors duration-200 hover:bg-sky-400"
      onClick={winner}
    >
      <CardContent className="text-4xl text-sky-950">{battlerName}</CardContent>
    </Card>
  );
}
