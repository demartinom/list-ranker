import { Separator } from "./ui/separator";
import { Popover, PopoverTrigger, PopoverContent } from "./ui/popover";
import HelpIcon from "@/assets/HelpIcon";

export default function Header() {
  return (
    <>
      {" "}
      <div className="mt-8 flex items-center justify-center text-center">
        <h1 className="text-5xl font-extrabold text-gray-700 md:text-7xl">
          List Ranker
        </h1>
        <Popover>
          <PopoverTrigger>
            <HelpIcon />
          </PopoverTrigger>
          <PopoverContent>Place content for the popover here.</PopoverContent>
        </Popover>
      </div>
      <Separator className="mt-5" style={{ height: "3px" }} />
    </>
  );
}
