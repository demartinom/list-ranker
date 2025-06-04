import Sun from "@/assets/Sun";
import Moon from "@/assets/Moon";
import { useTheme } from "@/components/theme-provider";
import { Button } from "./ui/button";

export default function ModeToggle() {
  const { theme, setTheme } = useTheme();
  const iconClasses = "size-6 md:size-10";

  const buttonIcon =
    theme == "light" ? (
      <Sun classname={iconClasses} />
    ) : (
      <Moon classname={iconClasses} />
    );
  const changeTheme = () => {
    setTheme(theme == "light" ? "dark" : "light");
  };
  return (
    <Button
      variant={"ghost"}
      onClick={changeTheme}
      className="absolute top-2 right-10 size-10 align-baseline md:size-16"
    >
      {buttonIcon}
    </Button>
  );
}
