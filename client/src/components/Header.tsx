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
          <PopoverContent className="w-2xl">
            A fun and simple game to help you rank your favorite things. <br />
            Start by choosing a premade list. <br />
            Once you pick a list, you’ll be shown two items at a time. Just
            choose the one you like more in each pair.
            <br /> Your choices will shape the ranking as you go. When you're
            done, you'll see a personalized ranked list based on all your
            decisions.
          </PopoverContent>
        </Popover>
      </div>
      <Separator className="mt-5" style={{ height: "3px" }} />
    </>
  );
}
