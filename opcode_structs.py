from dataclasses import dataclass
from typing import Literal

# https://hh.gbdev.io/


@dataclass(frozen=True)
class Operand:

    immediate: bool
    name: str
    bytes: int
    value: int | None
    adjust: Literal["+", "-"] | None

    def create(self, value):
        return Operand(
            immediate=self.immediate,
            name=self.name,
            bytes=self.bytes,
            value=value,
            adjust=self.adjust,
        )


@dataclass
class Instruction:

    opcode: int
    immediate: bool
    operands: list[Operand]
    cycles: list[int]
    bytes: int
    mnemonic: str
    comment: str = ""

    def create(self, operands):
        return Instruction(
            opcode=self.opcode,
            immediate=self.immediate,
            operands=operands,
            cycles=self.cycles,
            bytes=self.bytes,
            mnemonic=self.mnemonic,
        )


@dataclass
class Decoder:

    data: bytes
    address: int
    prefixed_instructions: dict
    instructions: dict

    @classmethod
    def create(cls, opcode_file: Path, data: bytes, address: int = 0):

        # Loads the opcodes from the opcode file
        prefixed, regular = load_opcodes(opcode_file)

        return cls (
            prefixed_instructions=prefixed,
            instructions=regular,
            data=data,
            address=address,
        )

if __name__ == "__main__":
    print("Hello world")
