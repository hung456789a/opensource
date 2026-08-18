package main

import "fmt"

type TrangThai string

const (
	MoCua   TrangThai = "Mở cửa"
	DongCua TrangThai = "Đóng cửa"
	BaoTri  TrangThai = "Bảo trì"
)

type MonAn struct {
	Ten string
	Gia int
}

type ThucDon struct {
	DanhSach []MonAn
}

type NhanVien struct {
	Ten    string
	ChucVu string
}

type BanAn struct {
	SoBan   int
	CoKhach bool
}

type NhaHang struct {
	Ten       string
	SoKhach   int
	TrangThai TrangThai
	ThucDon   ThucDon
	NhanVien  []NhanVien
	BanAn     []BanAn
}

// -------------------------------------------------------------
// 2. HÀM KHỞI TẠO (Constructor)
// Dấu * chỉ định: Kết quả trả về là một ĐỊA CHỈ, không phải dữ liệu
// -------------------------------------------------------------
func NewNhaHang(ten string) *NhaHang {

	// Dấu & nghĩa là: Tạo một nhà hàng thật trong RAM,
	// sau đó lấy địa chỉ (chìa khóa) của nó để trả về.
	return &NhaHang{
		Ten:       ten,
		SoKhach:   0,
		TrangThai: MoCua,
		ThucDon: ThucDon{
			DanhSach: []MonAn{
				{Ten: "Phở", Gia: 50000},
				{Ten: "Bún", Gia: 40000},
			},
		},
		NhanVien: []NhanVien{},
		BanAn:    []BanAn{},
	}
}

// -------------------------------------------------------------
// 3. METHOD (Phương thức - Đã học ở bài trước)
// Dùng *NhaHang (Pointer Receiver) để thay đổi số khách thật
// -------------------------------------------------------------
func (n *NhaHang) DonKhachMoi() {
	if n.TrangThai != MoCua {
		fmt.Printf("[%s] Nhà hàng đang %s, không thể đón khách!\n", n.Ten, n.TrangThai)
		return
	}
	n.SoKhach++
	fmt.Printf("[%s] Vừa đón thêm 1 khách. Tổng số khách: %d\n", n.Ten, n.SoKhach)
}

func (n *NhaHang) TuyenNhanVien() {
	nhanVien := NhanVien{}
	n.NhanVien = append(n.NhanVien, nhanVien)
	fmt.Printf("[%s] Vừa tuyển thêm 1 nhân viên. Tổng số nhân viên: %d\n", n.Ten, len(n.NhanVien))
}

func main() {
	// Bước 1: Gọi hàm khởi tạo.
	// Biến 'quanPho' lúc này không chứa nguyên cái nhà hàng, nó chỉ cầm CHÌA KHÓA.
	quanPho := NewNhaHang("Phở Bát Đàn")
	fmt.Printf("--- trạng thái ban đầu ---\n")
	fmt.Printf("Tên: %s, Số khách: %d, Trạng thái: %s\n", quanPho.Ten, quanPho.SoKhach, quanPho.TrangThai)
	fmt.Printf("Số món trong menu: %d món \n", len(quanPho.ThucDon.DanhSach))

	// Thử cho một nhân viên A cầm chìa khóa này
	nhanVienA := quanPho

	// Thử cho nhân viên B cầm chung chìa khóa (địa chỉ)
	nhanVienB := quanPho

	// Bước 2: Cả 2 nhân viên đều gọi hàm đón khách
	fmt.Println("--- Bắt đầu kinh doanh ---")
	nhanVienA.DonKhachMoi() // Khách thứ 1
	nhanVienB.DonKhachMoi() // Khách thứ 2
	quanPho.DonKhachMoi()   // Ông chủ (quanPho) tự đón khách thứ 3

	// Bước 3: Kiểm tra kết quả
	fmt.Println("--------------------------")
	fmt.Printf("Chốt sổ: %s hôm nay có %d khách.\n", quanPho.Ten, quanPho.SoKhach)
}
