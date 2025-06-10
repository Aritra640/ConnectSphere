import { atom } from "recoil";

type GroupMessageSelect = {
  Valid: boolean;
  GroupAvatar: string;
  GroupName: string;
  GroupId: string;
}

export const GroupMessageSelectAtom = atom<GroupMessageSelect> ({
  default: {
    Valid: false,
    GroupAvatar: "abcer",
    GroupName: "Unknown",
    GroupId: "1223",
  },
  key: "GroupMessageSelectAtom",
});
