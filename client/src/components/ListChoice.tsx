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
      className="m-2.5 h-22 w-fit min-w-40 cursor-pointer bg-sky-300 px-10 text-2xl text-sky-950 hover:bg-sky-500"
      onClick={listSelection}
    >
      {listName}
    </Button>
  );
}
