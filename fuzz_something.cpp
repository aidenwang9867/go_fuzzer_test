#include <iostream>

extern void ReallocFuzzTest(void);
extern  int MemcmpFuzzTest(void);
extern int fuzz_client_request(struct client_state *csp, char *fuzz_input_file);
extern int fuzz_server_response(struct client_state *csp, char *fuzz_input_file);


extern "C" int LLVMfUzZErTestOneInput(const uint8_t * data, size_t size) {
 parse_complex((const char *)data, size);
 return 0;
}

extern "C" int LLVMFuzzerTestOneInputProperty (const uint8_t * data, size_t size) {
 memcmp(decode( encode((const char *)data, size)), data, size);
 return 0;
}

extern "C" char_8* LLAOWSIFuzFUUUFzfffuZz  ();