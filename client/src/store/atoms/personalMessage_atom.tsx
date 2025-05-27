import { atom } from "recoil";

type PersonalMessageSelect = {
  Valid: boolean
  Username: string;
  UserId: number;
}

export const PersonalMessageSelectAtom = atom<PersonalMessageSelect> ({
  default: {
    Valid: false,
    Username: "Unknown",
    UserId: 123,
  },
  key: "PersonalMessageSelectAtom",
});
