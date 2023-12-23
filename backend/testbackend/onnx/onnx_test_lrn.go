package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("LRN", "TestLrn", NewTestLrn)
}

// NewTestLrn version: 3.
func NewTestLrn() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "LRN",
		Title:  "TestLrn",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x8f, 0x1, 0xa, 0x49, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x3, 0x4c, 0x52, 0x4e, 0x2a, 0xf, 0xa, 0x5, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x15, 0x17, 0xb7, 0x51, 0x39, 0xa0, 0x1, 0x1, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x65, 0x74, 0x61, 0x15, 0x0, 0x0, 0x0, 0x3f, 0xa0, 0x1, 0x1, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x15, 0x0, 0x0, 0x0, 0x40, 0xa0, 0x1, 0x1, 0x2a, 0xb, 0xa, 0x4, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x3, 0xa0, 0x1, 0x2, 0x12, 0x8, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x72, 0x6e, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "LRN",
		     Attributes: ([]*ir.AttributeProto) (len=4 cap=4) {
		    (*ir.AttributeProto)(0xc000133e00)(name:"alpha" type:FLOAT f:0.0002 ),
		    (*ir.AttributeProto)(0xc000133f00)(name:"beta" type:FLOAT f:0.5 ),
		    (*ir.AttributeProto)(0xc000126000)(name:"bias" type:FLOAT f:2 ),
		    (*ir.AttributeProto)(0xc000126100)(name:"size" type:INT i:3 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5, 5, 5, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067, 0.37642553, -1.0994008, 0.2982382, 1.3263859, -0.69456786, -0.14963454, -0.43515354, 1.8492638, 0.67229474, 0.40746182, -0.76991606, 0.5392492, -0.6743327, 0.031830557, -0.6358461, 0.67643327, 0.57659084, -0.20829876, 0.3960067, -1.0930616, -1.4912575, 0.4393917, 0.1666735, 0.63503146, 2.3831449, 0.94447947, -0.91282225, 1.1170163, -1.3159074, -0.4615846, -0.0682416, 1.7133427, -0.74475485, -0.82643855, -0.09845252, -0.6634783, 1.1266359, -1.0799315, -1.1474687, -0.43782005, -0.49803245, 1.929532, 0.9494208, 0.08755124, -1.2254355, 0.844363, -1.0002153, -1.5447711, 1.1880298, 0.3169426, 0.9208588, 0.31872764, 0.8568306, -0.6510256, -1.0342429, 0.6815945, -0.80340964, -0.6895498, -0.4555325, 0.017479159, -0.35399392, -1.3749512, -0.6436184, -2.2234032, 0.62523144, -1.6020577, -1.1043833, 0.05216508, -0.739563, 1.5430146, -1.2928569, 0.26705086, -0.039282817, -1.1680934, 0.5232767, -0.17154633, 0.77179056, 0.82350415, 2.163236, 1.336528, -0.36918184, -0.23937918, 1.0996596, 0.6552637, 0.64013153, -1.616956, -0.024326125, -0.7380309, 0.2799246, -0.09815039, 0.9101789, 0.3172182, 0.78632796, -0.4664191, -0.94444627, -0.4100497, -0.017020414, 0.37915173, 2.259309, -0.042257152, -0.955945, -0.34598178, -0.463596, 0.48148146, -1.540797, 0.06326199, 0.15650654, 0.23218104, -0.5973161, -0.23792173, -1.424061, -0.49331987, -0.54286146, 0.41605005, -1.1561824, 0.7811981, 1.4944845, -2.069985, 0.42625874, 0.676908, -0.63743705, -0.3972718, -0.13288058, -0.29779088, -0.30901298, -1.6760038, 1.1523316, 1.0796186, -0.81336427, -1.4664243, 0.5210649, -0.57578796, 0.14195317, -0.31932843, 0.69153875, 0.6947491, -0.7255974, -1.383364, -1.5829384, 0.6103794, -1.1888592, -0.5068163, -0.596314, -0.052567296, -1.9362798, 0.1887786, 0.52389103, 0.08842209, -0.31088617, 0.097400166, 0.39904633, -2.7725928, 1.9559124, 0.39009333, -0.6524086, -0.39095336, 0.49374178, -0.11610394, -2.0306845, 2.064493, -0.11054066, 1.0201727, -0.69204986, 1.5363771, 0.2863437, 0.60884386, -1.0452534, 1.2111453, 0.68981814, 1.3018463, -0.6280876, -0.48102713, 2.3039167, -1.0600158, -0.1359497, 1.1368914, 0.09772497, 0.5829537, -0.39944902, 0.37005588, -1.3065269, 1.6581306, -0.11816405, -0.6801782, 0.6663831, -0.4607198, -1.3342584, -1.3467175, 0.69377315, -0.15957344, -0.13370156, 1.0777438, -1.1268258, -0.7306777, -0.3848798, 0.09435159, -0.042171452, -0.2868872, -0.0616264, -0.10730527, -0.7196044, -0.812993, 0.27451634, -0.8909151, -1.1573553, -0.31229225, -0.15766701, 2.2567234, -0.7047003, 0.9432607, 0.7471883, -1.1889449, 0.77325296, -1.1838807, -2.6591723, 0.60631955, -1.7558906, 0.45093447, -0.6840109, 1.6595508, 1.0685093, -0.4533858, -0.6878376, -1.2140774, -0.44092262, -0.28035548, -0.36469355, 0.15670386, 0.5785215, 0.34965447, -0.76414394, -1.4377915, 1.3645319, -0.6894492, -0.6522936, -0.52118933, -1.8430696, -0.477974, -0.4796558, 0.6203583, 0.6984571, 0.003770889, 0.93184835, 0.339965, -0.015682112, 0.16092817, -0.19065349, -0.3948495, -0.26773354, -1.1280113, 0.2804417, -0.9931236, 0.8416313, -0.24945858, 0.04949498, 0.4938368, 0.6433145, -1.5706234, -0.20690368, 0.8801789, -1.6981058, 0.38728046, -2.2555642, -1.0225068, 0.038630553, -1.6567152, -0.98551077, -1.471835, 1.648135, 0.16422775, 0.5672903, -0.2226751, -0.35343176, -1.6164742, -0.29183736, -0.7614922, 0.8579239, 1.1411018, 1.4665787, 0.85255194, -0.5986539, -1.1158969, 0.7666632, 0.3562928, -1.7685385, 0.3554818, 0.8145198, 0.058925588, -0.18505368, -0.8076485, -1.4465348, 0.800298, -0.30911446, -0.23346666, 1.7327212, 0.6845011, 0.370825, 0.1420618, 1.5199949, 1.7195894, 0.9295051, 0.5822246, -2.094603, 0.12372191, -0.13010696, 0.09395323, 0.9430461, -2.7396772, -0.56931204, 0.26990435, -0.46684554, -1.4169061, 0.8689635, 0.27687192, -0.97110456, 0.3148172, 0.8215857, 0.005292646, 0.8005648, 0.078260176, -0.39522898, -1.1594205, -0.085930765, 0.19429293, 0.87583274, -0.11510747, 0.4574156, -0.964612, -0.78262913, -0.1103893, -1.0546285, 0.8202478, 0.46313033, 0.27909577, 0.3389041, 2.0210435, -0.4688642, -2.2014413, 0.1993002, -0.050603542, -0.51751906, -0.97882986, -0.43918952, 0.18133843, -0.5028167, 2.4124537, -0.96050435, -0.79311734, -2.28862, 0.25148442, -2.0164065, -0.53945464, -0.27567053, -0.70972794, 1.7388726, 0.99439436, 1.3191369, -0.8824188, 1.128594, 0.49600095, 0.77140594, 1.0294389, -0.90876323, -0.42431763, 0.86259604, -2.6556191, 1.5133281, 0.55313206, -0.045703962, 0.22050765, -1.0299352, -0.34994337, 1.1002843, 1.298022, 2.696224, -0.07392467, -0.65855294, -0.51423395, -1.0180418, -0.07785475, 0.38273242, -0.03424228, 1.0963469, -0.2342158, -0.34745064, -0.5812685, -1.6326345, -1.5677677, -1.179158, 1.3014281, 0.8952603, 1.3749641, -1.3322116, -1.9686247, -0.6600563, 0.17581895, 0.49869028, 1.0479722, 0.28427967, 1.7426687, -0.22260568, -0.9130792, -1.6812183, -0.8889713, 0.24211796, -0.8887203, 0.9367425, 1.4123276, -2.369587, 0.8640523, -2.239604, 0.40149906, 1.2248706, 0.064856105, -1.2796892, -0.5854312, -0.26164544, -0.18224478, -0.20289683, -0.10988278, 0.21348006, -1.2085737, -0.24201983, 1.5182612, -0.38464543, -0.4438361, 1.0781974, -2.5591846, 1.1813786, -0.63190377, 0.16392857, 0.09632136, 0.9424681, -0.26759475, -0.6780258, 1.2978458, -2.364174, 0.020334182, -1.3479254, -0.7615734, 2.0112567, -0.044595428, 0.1950697, -1.7815628, -0.7290447, 0.1965574, 0.3547577, 0.61688656, 0.008627899, 0.5270042, 0.4537819, -1.8297404, 0.037005723, 0.76790243, 0.5898798, -0.36385882, -0.8056265, -1.1183119, -0.13105401, 1.1330799, -1.9518042, -0.6598917, -1.1398025, 0.7849575, -0.5543096, -0.47063765, -0.21694957, 0.44539326, -0.392389, -3.046143, 0.5433119, 0.43904296, -0.21954103, -1.0840366, 0.35178012, 0.37923554, -0.47003287, -0.21673147, -0.9301565, -0.17858909, -1.5504293, 0.41731882, -0.9443685, 0.23810315, -1.405963, -0.5900577, -0.110489406, -1.6606998, 0.115147874, -0.37914756, -1.7423562, -1.3032428, 0.60512006, 0.895556, -0.13190864, 0.40476182, 0.22384356, 0.32962298, 1.285984, -1.5069984, 0.67646074, -0.38200897, -0.22425893, -0.30224973, -0.3751471}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5, 5, 5, 5),
				tensor.WithBacking([]float32{1.2472647, 0.2829531, 0.6920608, 1.5843562, 1.3204386, -0.69102854, 0.6718023, -0.10702425, -0.07298197, 0.2903356, 0.10185411, 1.0282644, 0.5381168, 0.08603699, 0.3138572, 0.23593862, 1.0563987, -0.14506166, 0.2213579, -0.60392696, -1.8050345, 0.46216273, 0.61123496, -0.52476233, 1.6048199, -1.0282886, 0.032355987, -0.13235621, 1.0836803, 1.0388957, 0.109562255, 0.2673966, -0.6277499, -1.4005333, -0.2460093, 0.11055451, 0.8698924, 0.8501729, -0.27386707, -0.2137592, -0.74142194, -1.0039887, -1.2064515, 1.3792979, -0.36037257, -0.3097279, -0.8858304, 0.5497444, -1.141109, -0.15041667, -0.63315475, 0.27357724, -0.36118996, -0.8347796, -0.019927127, 0.30286846, 0.047034025, 0.21387602, -0.4484835, -0.2564862, -0.47546908, -0.25422964, -0.57496077, -1.2205791, 0.12545665, -0.2840953, -1.1526058, 0.3272192, -0.64149714, 0.036730703, 0.5155344, 0.09120228, 0.8056096, -0.8730934, 0.28449675, -0.48419496, -0.6157186, -0.40929347, -0.22029197, 0.03971398, -0.8238123, 0.6369701, 0.32926834, -1.0861714, 1.0522736, 1.3404278, 0.83349055, -0.12722303, -0.75703835, 0.7455948, -0.28508455, 0.86432636, 0.14727196, 0.6905633, 0.25198498, 0.49961293, 0.0074244835, 1.2627048, 0.08973546, 0.2842456, 1.3314996, -0.95296866, -0.8983394, 0.68545514, -0.82950443, 1.37423, -0.29246798, -0.5285235, 1.359588, 1.0468051, 1.3204077, 0.64064676, -0.60897064, 1.350512, -0.18950327, 0.5674147, 0.6697816, -0.10960856, 0.43421006, 0.652088, 0.2661702, -0.7773781, 0.2108747, 0.9378687, -0.49112839, -0.10580754, -0.307684, 1.3075403, 0.4753752, 0.28811818, -0.5444035, 0.38129687, -0.4768123, 0.022507109, -0.4496066, 0.47830492, 0.40768373, -0.14728715, 0.28001824, -0.77287656, -1.0544267, 0.31069067, 0.11785122, 0.44902146, 1.6849756, 0.66782844, -0.64545274, 0.7898237, -0.93045354, -0.3263826, -0.048253708, 1.2114402, -0.5265822, -0.5843673, -0.06961624, -0.46914098, 0.79660606, -0.76360095, -0.81129825, -0.30958045, -0.35214293, 1.3642653, 0.6713313, 0.061907344, -0.8664404, 0.5970089, -0.70724404, -1.0922742, 0.8400194, 0.22408965, 0.6511263, 0.22536872, 0.6058439, -0.4602922, -0.7312828, 0.48195523, -0.568062, -0.48756716, -0.32210302, 0.012359546, -0.25029823, -0.97218615, -0.45509082, -1.5720174, 0.442101, -1.1327571, -0.78085136, 0.036885347, -0.52294326, 1.0909894, -0.914149, 0.1888301, -0.027775975, -0.8258583, 0.3700102, -0.121297956, 0.54573095, 0.5822896, 1.5295027, 0.9449856, -0.26104835, -0.16926458, 0.7775542, 0.46333373, 0.45263785, -1.1432697, -0.017200556, -0.5218558, 0.19791944, -0.069400795, 0.6435507, 0.22429386, 0.5559724, -0.3298029, -0.66778284, -0.2899381, -0.012035204, 0.26810002, 1.5973982, -0.029880133, -0.6759129, -0.24463773, -0.32780063, 0.3404272, -1.0893934, 0.04473288, 0.110666685, 0.16417333, -0.42236072, -0.16823477, -1.0068853, -0.3488284, -0.38385567, 0.2941906, -0.8175261, 0.5523772, 1.0567191, -1.4635807, 0.30140844, 0.4786355, -0.45073172, -0.28091288, -0.09396051, -0.21055172, -0.21850483, -1.1850401, 0.8148018, 0.76338804, -0.5751268, -1.0368404, 0.3684459, -0.40713966, 0.100375995, -0.2257834, 0.48895314, 0.49125776, -0.5130614, -0.97814703, -1.1192156, 0.43160015, -0.8406254, -0.3583652, -0.42164487, -0.037170395, -1.3690324, 0.13348567, 0.3704438, 0.06251832, -0.21982525, 0.06887229, 0.28216153, -1.9602677, 1.3829428, 0.2758362, -0.46131822, -0.27643594, 0.34910882, -0.08209782, -1.4357986, 1.4596908, -0.078163125, 0.7213308, -0.48931885, 1.0862858, 0.20247395, 0.43050468, -0.7390748, 0.8563649, 0.48776686, 0.9204585, -0.44412175, -0.34013465, 1.6289686, -0.74952906, -0.09613089, 0.8038772, 0.06909236, 0.41218138, -0.2824479, 0.26166072, -0.9238239, 1.1724164, -0.083547466, -0.4809179, 0.47116005, -0.32577384, -0.9433966, -0.9522272, 0.490537, -0.11282197, -0.0945401, 0.7620122, -0.79674715, -0.5166545, -0.27213028, 0.06671493, -0.029819502, -0.20284006, -0.04357456, -0.075876, -0.50882107, -0.57486516, 0.19411094, -0.62995857, -0.818352, -0.22081517, -0.11147841, 1.5955591, -0.49828643, 0.6669665, 0.52833277, -0.8406187, 0.54674834, -0.83710057, -1.8800844, 0.4287264, -1.2415143, 0.31884637, -0.48365977, 1.1734228, 0.75553536, -0.32059088, -0.48636886, -0.8584602, -0.3117717, -0.19823903, -0.25786963, 0.11080487, 0.40906835, 0.24723703, -0.5403252, -1.0166366, 0.9647579, -0.48750627, -0.46123114, -0.36853144, -1.3031425, -0.33797398, -0.33915865, 0.43860507, 0.49387673, 0.002666284, 0.6589045, 0.24038923, -0.011088419, 0.11379119, -0.13481183, -0.27919784, -0.18931133, -0.797605, 0.19830172, -0.7022313, 0.59511596, -0.17639269, 0.034998164, 0.34919184, 0.45487946, -1.1105419, -0.1463026, 0.62234, -1.2006818, 0.27384493, -1.5947894, -0.7230085, 0.027315626, -1.17138, -0.6968426, -1.0407053, 1.1653535, 0.116120696, 0.40112954, -0.15745458, -0.2499134, -1.142926, -0.20634972, -0.5384433, 0.60663295, 0.80680436, 1.0369903, 0.6028378, -0.42330793, -0.7890248, 0.54201734, 0.25193498, -1.2504628, 0.25135005, 0.5759254, 0.041662622, -0.13085018, -0.57107854, -1.0227704, 0.5658746, -0.2185687, -0.16507646, 1.2251569, 0.48400772, 0.26220617, 0.1004526, 1.0747098, 1.2158562, 0.65724343, 0.41168615, -1.4809446, 0.08748055, -0.09199835, 0.066433325, 0.6668031, -1.9369761, -0.40256086, 0.19084065, -0.33008528, -1.0018556, 0.6143926, 0.19577752, -0.68665636, 0.22260024, 0.5809268, 0.003742448, 0.56607795, 0.055335294, -0.2794391, -0.8198013, -0.060761563, 0.13736849, 0.6192682, -0.08138657, 0.32343715, -0.6820224, -0.55339193, -0.07805305, -0.74570876, 0.5799709, 0.32743624, 0.19734499, 0.23963967, 1.428977, -0.3315189, -1.5564873, 0.14092582, -0.0357811, -0.365896, -0.692092, -0.31055132, 0.1282242, -0.35554326, 1.7056623, -0.6791521, -0.5608014, -1.618111, 0.17780234, -1.4257177, -0.3814461, -0.19492438, -0.5018354, 1.2295063, 0.7031184, 0.9327331, -0.623954, 0.7980186, 0.3507235, 0.5454238, 0.72790766, -0.6425319, -0.3000368, 0.60993993, -1.8775772, 1.0700266, 0.3911202, -0.032317564, 0.15592167, -0.72819066, -0.24744302, 0.77799463, 0.91773427, 1.9062853, -0.05226909, -0.46566162, -0.36361626, -0.7198458, -0.055048846, 0.27062848, -0.024212593, 0.77519304, -0.16559993, -0.24568115, -0.41098222, -1.1543926, -1.1085061, -0.8337712, 0.92019755, 0.63303256, 0.9722147, -0.94198745, -1.3919371, -0.46672678, 0.124322616, 0.35261723, 0.74101394, 0.2010081, 1.2321875, -0.15740533, -0.645623, -1.1886151, -0.6285748, 0.17120194, -0.6284101, 0.66236717, 0.99859846, -1.6753907, 0.6109637, -1.5834534, 0.2838629, 0.86605716, 0.04585774, -0.9048179, -0.4139266, -0.18500522, -0.12886256, -0.14345278, -0.07769759, 0.1509529, -0.85456455, -0.17112948, 1.0735301, -0.2719697, -0.31383714, 0.76233274, -1.809334, 0.8353222, -0.44681746, 0.11591317, 0.06810774, 0.66637963, -0.18920007, -0.47941676, 0.9175549, -1.6715511, 0.014377767, -0.9530885, -0.5384911, 1.422064, -0.031533666, 0.13793449, -1.2596844, -0.51542807, 0.1389862, 0.25084412, 0.43620116, 0.006100492, 0.3726448, 0.32086936, -1.2937198, 0.02616412, 0.5429632, 0.4171026, -0.25727606, -0.5696561, -0.790726, -0.09266896, 0.8011588, -1.3799995, -0.46656695, -0.8059075, 0.5550262, -0.39194936, -0.3327506, -0.15340205, 0.31493744, -0.27744183, -2.1535957, 0.38417634, 0.31044835, -0.15523757, -0.76649344, 0.24873504, 0.2681564, -0.33234283, -0.15325204, -0.65770304, -0.12628046, -1.0962727, 0.29508492, -0.66774553, 0.16836414, -0.9941119, -0.41720486, -0.07812722, -1.1742128, 0.08142099, -0.2680958, -1.2319651, -0.92150503, 0.42788047, 0.6332436, -0.093259044, 0.28620765, 0.15828066, 0.23307805, 0.9092851, -1.0655663, 0.4783252, -0.2701195, -0.15857476, -0.21371943, -0.2652683}),
			),
		},
	}
}
