import { atom } from "recoil";

export const authmodal = atom<true|false>({
  default: false,
  key: "authmodal",
});
