import { Button } from "@/components/ui/button";

interface ListProps {
  listName: string;
  listSelection: () => Promise<void>;
}

export default function ListChoice({
  listName = "list",
  listSelection,
}: ListProps) {
  return (
    <Button variant={"outline"} onClick={listSelection}>
      {listName}
    </Button>
  );
}
