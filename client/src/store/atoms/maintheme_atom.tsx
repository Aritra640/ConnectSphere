import { atom } from "recoil";

export const MainThemeAtom = atom<"Bright" | "Dark">({
  default: "Dark",
  key: "MainThemeAtom",
});
