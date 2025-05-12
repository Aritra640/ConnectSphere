import { atom } from "recoil";

export const OTPModalAtom = atom<true | false>({
  default: false,
  key: "OTPModalAtom",
});
