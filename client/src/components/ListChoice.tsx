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
    <Button
      variant={"ghost"}
      className="choice-button m-2.5 h-24 w-fit min-w-[15rem] px-10 text-2xl"
      onClick={listSelection}
    >
      {listName}
    </Button>
  );
}
