import base64
import datetime

def decode_short_request_id(request_id):
    if len(request_id) != 16:
        raise ValueError(f"Invalid Request ID length: {len(request_id)} (expected 16)")
    
    request_id = request_id.upper()
    
    try:
        decoded_bytes = base64.b32decode(request_id)
    except base64.binascii.Error as e:
        raise ValueError(f"Base32 decoding failed: {e}")
    
    if len(decoded_bytes) != 10:
        raise ValueError(f"Decoded byte length mismatch: {len(decoded_bytes)} (expected 10)")
    
    timestamp_bytes = decoded_bytes[:8]
    
    timestamp = int.from_bytes(timestamp_bytes, byteorder="big", signed=True)
    
    random_bytes = decoded_bytes[8:]
    random_int = int.from_bytes(random_bytes, byteorder="big")
    
    dt = datetime.datetime.fromtimestamp(timestamp / 1000.0, tz=datetime.timezone.utc)
    
    return {
        "raw_id": request_id,
        "bytes": decoded_bytes.hex(),
        "timestamp": timestamp,
        "utc_time": dt,
        "random_bytes": random_bytes.hex(),
        "random_int": random_int
    }

if __name__ == "__main__":
    sample_id = input("Request ID: ")
    
    try:
        decoded = decode_short_request_id(sample_id)
        
        print("Result:")
        print(f"    Original Request ID: {decoded['raw_id']}")
        print(f"    Byte Request ID: {decoded['bytes']}")
        print(f"    Timestamp: {decoded['timestamp']} ms")
        print(f"    UTC Time: {decoded['utc_time'].isoformat()}Z")
        print(f"    Random Data Byte: {decoded['random_bytes']}")
        print(f"    Random Data Int: {decoded['random_int']}")
        
    except ValueError as e:
        print(f"Failed to decode: {e}")