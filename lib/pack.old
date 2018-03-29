package cmd

import (
  "github.com/spf13/cobra"
  "os"
  "fmt"
  "strings"
  "strconv"
)

type MetaImage struct {
  library, id, description, author, binfile string
  major, minor, bugfix int
}

var packCmd = &cobra.Command{
    Use:   "pack",
    Short: "Make a pack-bitstream out of multiple image-bitstreams",
    Long: `A set of bitstream images are appended and an index is generated. Pointers for five of them are added to the first positions of the resulting pack. This resemble the 'applet' expected by iCE40 devices when cold/warm boot is used.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
        fmt.Println(cmd)
        fmt.Println(len(args))
        for

        //fmt.Println(RootCmd.PersistentFlags(verbose))
        /*
        int main(int argc, char **argv) {
            Header headers[NUM_HEADERS];
            std::unique_ptr<Image> images[NUM_IMAGES];
            const char *outfile_name = NULL;

            for (int i = 1; i < argc; i++)


            if (!image_count)
                usage();

            if (coldboot && por_image != 0)
                error("Can't select power on reset boot image in cold boot mode\n");

            if (por_image >= image_count)
                error("Specified non-existing image for power on reset\n");

            // Place images
            uint32_t offs = NUM_HEADERS * HEADER_SIZE;
            if (align_first)
                align_offset(offs, align_bits);
            for (int i=0; i<image_count; i++) {
                images[i]->place(offs);
                offs += images[i]->size();
                align_offset(offs, align_bits);
                info("Place image %d at %06x .. %06x.\n", i, int(images[i]->offset()), int(offs));
            }

            // Populate headers
            for (int i=0; i<image_count; i++)
                headers[i + 1] = Header(*images[i]);
            headers[0] = headers[por_image + 1];
            for (int i=image_count; i < NUM_IMAGES; i++)
                headers[i + 1] = headers[0];
            if (coldboot)
                headers[0].set_coldboot_flag();


            uint32_t file_offset = 0;
            for (int i=0; i<NUM_HEADERS; i++)
            {
                pad_to(*osp, file_offset, i * HEADER_SIZE);
                headers[i].write(*osp, file_offset);
            }
            for (int i=0; i<image_count; i++)
            {
                pad_to(*osp, file_offset, images[i]->offset());
                images[i]->write(*osp, file_offset);
            }
        */
    },
}

var Verbose int
var Mode bool
var Compress string
var Output string
var Base int
var PointersStr string
var Pointers [4]int
var Align int

func init() {
  RootCmd.AddCommand(packCmd)
  packCmd.PersistentFlags().IntVarP(&Verbose,"verbose", "v", 0, "0: disabled, 1: minimum, 2: full")
  packCmd.PersistentFlags().BoolVarP(&Mode,"m", "m", false, "enable cold/warm boot mode (pointer selected by CBSEL0:1)")
  packCmd.PersistentFlags().StringVarP(&Output,"output", "o", "", "write output pack-bitstream to file instead of stdout")
  packCmd.PersistentFlags().IntVarP(&Base,"base", "b", 0, "power on reset image when not using cold/warm boot")
  packCmd.PersistentFlags().StringVarP(&PointersStr,"pointers", "p", "1,2,3,4", "four default-addressable pointers")
  packCmd.PersistentFlags().IntVarP(&Align,"align", "a", 0, "align images at 2^<a> bytes")

  split := strings.Split(PointersStr,",")
  for i, v := range split {
    a, err := strconv.Atoi(v)
    if err != nil { os.Exit(1); }
    Pointers[i] = a
  }
  if len(split)>3 {
    fmt.Printf("WARNING: 'pointers' contains too many indexes. Since iCE40 devices support up to four, only %d are used.\n", Pointers)
  }
}

/*
class Header {
    uint32_t image_offs;
    bool coldboot_flag;
    bool empty;
public:
    Header() : empty(true) {}
    Header(const Image &i) :
        image_offs(i.offset()), coldboot_flag(false), empty(false) {}
    void set_coldboot_flag() { coldboot_flag = true; }
    void write(std::ostream &ofs, uint32_t &file_offset);
};
*/
type Header struct {
  image_offs int
  coldboot_flag, empty bool
}

func (h *Header) Header(f string) {}
func (h *Header) SetColdbootFlag(f string) { h.coldboot_flag = true }
func (h *Header) Write(stream *[]byte, pointer int) {
/*
Header::write

// Preamble
0x7eaa997e
// Boot mode
0x9200 (coldboot_flag? 0x10: 0x00)
// Boot address
0x4403
(image_offs >> 16) & 0xff
(image_offs >> 8) & 0xff
image_offs & 0xff
// Bank offset
0x820000
// Reboot
0x0108

// Zero out any unused bytes
while (file_offset & (HEADER_SIZE - 1))
    write_byte(ofs, file_offset, 0x00);
*/
}

//----

/*
-a<n>, -A<n>
align images at 2^<n> bytes. -A also aligns image 0.

static const int NUM_IMAGES = 4;
static const int NUM_HEADERS = NUM_IMAGES + 1;
static const int HEADER_SIZE = 32;

static void align_offset(uint32_t &offset, int bits)
{
    uint32_t mask = (1 << bits) - 1;
    if (offset & mask)
        offset = (offset | mask) + 1;
}

static void write_byte(std::ostream &ofs, uint32_t &file_offset, uint8_t byte)
{
    ofs << byte;
    file_offset++;
}

static void write_bytes(std::ostream &ofs, uint32_t &file_offset,
                        const uint8_t *buf, size_t n)
{
    if (n > 0) {
        ofs.write(reinterpret_cast<const char*>(buf), n);
        file_offset += n;
    }
}

static void write_file(std::ostream &ofs, uint32_t &file_offset,
                       std::istream &ifs)
{
    const size_t bufsize = 8192;
    uint8_t *buffer = new uint8_t[bufsize];

    while(!ifs.eof()) {
        ifs.read(reinterpret_cast<char *>(buffer), bufsize);
        if (ifs.bad())
            error("Read error on input image");
        write_bytes(ofs, file_offset, buffer, ifs.gcount());
    }

    delete[] buffer;
}

static void pad_to(std::ostream &ofs, uint32_t &file_offset, uint32_t target)
{
    if (target < file_offset)
        error("Trying to pad backwards!\n");
    while(file_offset < target)
        write_byte(ofs, file_offset, 0xff);
}

class Image {
    std::ifstream ifs;
    uint32_t offs;

public:
    Image(const char *filename) : ifs(filename, std::ifstream::binary) {}

    size_t size();
    void write(std::ostream &ofs, uint32_t &file_offset);
    void place(uint32_t o) { offs = o; }
    uint32_t offset() const { return offs; }
};

size_t Image::size()
{
    ifs.seekg (0, ifs.end);
    size_t length = ifs.tellg();
    ifs.seekg (0, ifs.beg);
    return length;
}

void Image::write(std::ostream &ofs, uint32_t &file_offset)
{
    write_file(ofs, file_offset, ifs);
}
*/
