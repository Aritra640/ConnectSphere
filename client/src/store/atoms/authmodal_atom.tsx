import { atom } from "recoil";

export const AuthModalAtom = atom<true|false>({
  default: false,
  key: "AuthModalAtom",
});
