import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const customListToArray = (list: string) => {
  return list.split("\n");
};

export const listValidation = (list: string): string | null => {
  const arrayFromList = customListToArray(list);
  if (arrayFromList.length < 4) {
    return "Please submit a list of four items or more.";
  }

  if (new Set(arrayFromList).size !== arrayFromList.length) {
    return "Duplicates found in list.";
  }

  if (arrayFromList.length > 40) {
    return "Your list is too long and will take too much time to rank.";
  }
  return null;
};
