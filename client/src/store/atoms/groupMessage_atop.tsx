import { atom } from "recoil";

type GroupMessageSelect = {
  Valid: boolean;
  GroupName: string;
  GroupId: string;
}

export const GroupMessageSelectAtom = atom<GroupMessageSelect> ({
  default: {
    Valid: false,
    GroupName: "Unknown",
    GroupId: "1223",
  },
  key: "GroupMessageSelectAtom",
});
