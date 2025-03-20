import { Card, CardContent } from "@/components/ui/card";

export default function Battler({ name = "Item" }) {
  return (
    <Card className="m-2.5 h-60 w-2xl justify-center bg-sky-200 text-center shadow-sky-200">
      <CardContent className="text-4xl text-sky-900">{name}</CardContent>
    </Card>
  );
}
