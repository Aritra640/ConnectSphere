import { atom } from "recoil";


export const TestAtom = atom<true|false>({
  default: false,
  key: "TestAtom",
});
