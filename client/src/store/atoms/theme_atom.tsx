import { atom } from "recoil";

export const ThemeAtom = atom<"Bright"|"Dark">({
  default: "Dark",
  key: "ThemeAtom",
});
