import { Button } from "@/components/ui/button";
import { Skeleton } from "./ui/skeleton";

interface ListProps {
  listName: string;
  listSelection: () => Promise<void>;
  loading: boolean;
}
export default function ListChoice({
  listName = "list",
  listSelection,
  loading,
}: ListProps) {
  return loading ? (
    <Skeleton className="m-2.5 h-22 w-fit min-w-50 bg-gray-200" />
  ) : (
    <Button
      variant={"ghost"}
      className="m-2.5 h-22 w-fit min-w-[15rem] cursor-pointer bg-sky-300 px-10 text-2xl text-sky-950 hover:bg-sky-500 md:min-w-[15rem]"
      onClick={listSelection}
    >
      {listName}
    </Button>
  );
}
