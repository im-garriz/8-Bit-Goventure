import sys
import hypothesis.strategies as st
from hypothesis import given
from fields import FIELDS
import struct
from collections import namedtuple
from pathlib import Path

HEADER_START = 0x100
HEADER_END = 0x14F

# Header size as measured from the last element to the first + 1
HEADER_SIZE = (HEADER_END - HEADER_START) + 1

CARTRIDGE_HEADER = "".join(format_type for _, format_type in FIELDS)

CartridgeMetadata = namedtuple(
    "CartridgeMetadata",
    [field_name for field_name, _ in FIELDS if field_name is not None],
)


def read_cartridge_metadata(buffer, offset: int = 0x100):
    """
    Unpacks the cartridge metatdata from 'buffer' at 'offset' and
    return a 'CartridgeMetadata' object
    """
    data = struct.unpack_from(CARTRIDGE_HEADER, buffer, offset=offset)
    return CartridgeMetadata._make(data)


@given(
    data=st.binary(
        min_size=HEADER_SIZE + HEADER_START, max_size=HEADER_SIZE + HEADER_START
    )
)
def test_read_cartridge_metadata_smoketest(data):
    def read(offset, count=1):
        return data[offset : offset + count + 1]

    metadata = read_cartridge_metadata(data)
    assert metadata.title == read(0x134, 14)
    checksum = read(0x14E, 2)

    # The checksum is _big endian_ -- so we need to tell Python
    # to read it back in properly!
    assert metadata.global_checksum == int.from_bytes(checksum, sys.byteorder)


if __name__ == "__main__":
    p = Path("snake.gb")
    print(read_cartridge_metadata(p.read_bytes()))
