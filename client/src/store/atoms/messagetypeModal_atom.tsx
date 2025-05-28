import { atom } from "recoil";

export const MessageTypeModalAtom = atom<true|false>({
  default: false,
  key: "MessageTypeModalAtom",
});
