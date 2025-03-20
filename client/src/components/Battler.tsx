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
      className="m-2.5 h-60 w-2xl justify-center bg-sky-200 text-center shadow-sky-200 hover:cursor-pointer"
      onClick={winner}
    >
      <CardContent className="text-4xl text-sky-900">{battlerName}</CardContent>
    </Card>
  );
}
