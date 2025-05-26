import { atom } from "recoil";

export const DashBoardAtom = atom<"Home"|"Message"|"Group"|"Setting">({
  default: "Home",
  key: "DashBoardAtom",
});
