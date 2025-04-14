import { atom } from "recoil";

export const testAtom = atom<number>({
  default: 0,
  key: "testAtom",
});
