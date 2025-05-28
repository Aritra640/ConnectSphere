import { useRecoilState, useSetRecoilState } from "recoil";
import { MessageTypeModalAtom } from "@/store/atoms/messagetypeModal_atom";
import { Button } from "../ui/button";
import { PlusIcon } from "@/icons/plus_icon";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "../ui/dialog";

export function MessageTypeModalButton() {
  const setState = useSetRecoilState(MessageTypeModalAtom);

  return (
    <Button
      onClick={() => setState(true)}
      className="bg-transparent cursor-pointer hover:bg-gray-50/15"
      variant="ghost"
      size="icon"
    >
      <PlusIcon />
    </Button>
  );
}

export function MessageTypeModal() {
  const [open, setOpen] = useRecoilState(MessageTypeModalAtom);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Select Message Type</DialogTitle>
          <DialogDescription>
            Choose whether you want to send a text, image, or other type of message.
          </DialogDescription>
        </DialogHeader>

        {/* Modal Body - customize with message type options */}
        <div className="flex gap-3 py-4">
          <Button variant="outline">Text</Button>
          <Button variant="outline">Image</Button>
        </div>

        <DialogClose asChild>
          <Button variant="secondary">Close</Button>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
}
