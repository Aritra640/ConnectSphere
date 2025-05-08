import { atom } from "recoil";

export const HomeMenuAtom = atom<true|false>({
  default: false,
  key: "HomeThemeAtom",
});
