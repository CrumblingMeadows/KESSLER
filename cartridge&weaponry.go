package main

type cartridgeT struct{
	cartridge
	description
	calibre
	cartridgelength
	cartridgemass_grams float64
	projectilediameter float64
	projectilelength float64
	projectilemass float64
	muzzlevelocity_meterspersec float64
	recoil_indegrees float64
	spread_radians float64
	projectiles int
	warheadtype string
	kineticenergyKJ float64
//	subprojectile bool
	projectileID 
	subprojectileID 
}
///couldnt even finish working on the firearm & weapon platform assets even though i had months to finish this, it's so joever
type weapon struct{
	weaponID
	capableofbeingattachment bool
	name
	cartridge
	description
	calibre
	cartridgelength
	mass float64 ///stats such as mass, firerate, mag capcity, etc have IDs for the user to be able to customize them - if i even get around to doing that.
	firerate float64
	jammingpercentage float64
	recoilreduction float64 //can allow for customization to reduce or enhance recoil
	ergonomic float64
	auditoryrange int
	sightID
	muzzleattachID
	stockID
	magID
	conv
}
type attachment struct {
	weaponcustomization weapon
	projectileconv cartridge
}
var(///33/34/35 Hollowpointsubproj
///////writing this has made me realize how much of a fucking mess the logistics are for infantry ammunition for Nova Industries
	///---cartridges---
//10x25
	PM_C_GP_10_25_HP=cartridgeT{
		"10x25",
		"Standard pistol cartridge, only great as a last resort due to the relatively low kinetic energy output and short effective range.",
		10,
		25,
		35,
		7.62,
		12.7,
		28.3,
		2555,
		12.7,
		.0000234,
		1,
		"10x25 Hollow Point - PM:C:GP:10_25:HPe:2053,E23",
		88,
		1,
		34,
	}

	PM_C_GP_10_25_APFSS=cartridgeT{
		"10x25",
		"Standard anti-armor pistol cartridge, fires a high-velocity sabot with greater range than the standard 10x25 cartridge.",
		10,
		25,
		88,
		7.62,
		15,
		76.2,
		3100,
		13,
		.0000234,
		1,
		"10x25 Armor-Piercing Fin-Stabilized Sabot - PM:C:GP:10_25:APFSS:2077:E27",
		355,
		30,
		31,
	}
///23x54R Flechette Cartridges
	PM_C_CQB_Flechette4=cartridgeT{
		"23x54R",
		"Heavy Cartridge for the CQB N821 family of shotguns - 4 high output sabots are utilized to penetrate through light to medium kinetic armor, and easily shred through particle & EM shielding on infantry.",
		23,
		54,
		237,
		7.5,
		27,
		45.3,
		2297.37,
		37,
		.12,
		4,
		"WTi[YX] 4-Sabot Flechette - PM:C:CQB:23_54R:4WTi[YX]:7.5_27:2112:X74",
		118.7,
		2,
		0,
	}
	//note: flechette point detonate should initially count as 1 projectile before being erased and split off into its sub projectiles
	PM_C_CQB_Flechette4PD=cartridgeT{
		"23x54R",
		"Heavy Cartridge for the CQB N821 family of shotguns - 4 high output sabots are utilized to penetrate through light to medium kinetic armor, and easily shred through particle & EM shielding on infantry.\nThis is the the more favored Point-Detonate variant, in which the sabots seperate from the main projectile centimeters from their target set through an EM fuze in the main projectile casing to launch the sabots - with said fuze being set through integrated hardware at the end of the firearm's muzzle.\nAmazing for ranged combat - capable of excelling both at CQB and medium-range firefights.",
		23,
		54,
		257,
		7.5,
		27,
		45.3,
		2287.57,
		32,
		.007,
		1,
		"WTi[YX] 4-Sabot Flechette: Point-Deonate - PM:C:CQB:23_54R:4WTi[YX]:7.5_27:2113:X73-PD",
		119.37,
		3,
		2,
	}
	PM_C_CQB_Flechette12=cartridgeT{
		"23x54R",
		"Standard Cartridge for the CQB N821 family of shotguns - Loaded with 12 sabots seperating after launch capable of punching through light body-armor and tearing personnel particle/EM shielding apart.",
		23,
		54,
		97,
		5,
		15,
		7.1,
		2337.762,
		12.7,
		.155,
		12,
		"WTi[YX] 12-Sabot Flechette PM:C:CQB:23_54R:12:WTi[YX]:5_15:2100:Y21",
		19.27,
		4,
		0,
	}
	PM_C_CQB_Flechette12PD=cartridgeT{
		"23x54R",
		"Standard Cartridge for the CQB N821 family of shotguns - Loaded with 12 sabots seperating centimeters away a target with the use of an EM fuze set by integrated hardware at the muzzle of a compatible firearm - capable of punching through light body-armor and tearing personnel particle/EM shielding apart.",
		23,
		54,
		105.4,
		5,
		15,
		7.1,
		2332.724,
		11.3,
		.023,
		1,
		"WTi[YX] 12-Sabot Flechette PM:C:CQB:23_54R:12:WTi[YX]:5_15:2100:Y21-PD",
		19.33,
		5,
		4,
	}
	//23x54R Breaching (HEDP) - & EMP
	PM_C_CQB_HEDP=cartridgeT{
		"23x54R",
		"Standard breaching cartridge for the CQB N821 family of shotguns - Loaded with tandem shaped-charge warheads with a tertiary high-explosive charge in a single projectile with a point-detonate electromagnetic fuze, amazing for punching through particle/EM shielding.",
		23,
		54,
		57,
		16,
		33,
		37,
		2092,
		10.23,
		.05,
		1,
		"High-Explosive Dual-Purpose - PM:C:CQB:23_54R:12:HEDP:16_33:2099:Y23-PD",
		75,
		6,
		7,
	}
	PM_C_CQB_EMP=cartridgeT{
		"23x54R",
		"Specialized EMP cartridge for the CQB N821 family of shotguns - great for disabling unprotected hardware, disabling exposed particle and electromagnetic projectors, and inducing electrical burns on flesh lacking EMP protection.\nUseless against active particle and electromagnetic shielding, along with EMP resistant protection.",
		23,
		54,
		55,
		16,
		33,
		32,
		2037,
		10.23,
		.05,
		1,
		"EMP Warhead - PM:C:CQB:23_54R:EMP:16_33:2095:V37",
		64,
		8,
		9,
	}
	//9x54R Cartridges
	PM_C_GP_9x54R_HPe=cartridgeT{
		"9x54R",
		"High-output gauss cartridge for the standard N310 and previous 9x54R compatible firearms - kinetic expanding warhead similar to ancient hollow-point ammunition poses lethal damage against soft target as chunks of material embed themselves within flesh, light body-armor, and exposed hardware.",
		9,
		54,
		66,
		7.62,
		30,
		67,
		3557,
		17.8,
		.0023,
		1,
		"Kinetic Expansion (Hollow-Point Equivalent) - PM:C:GP:9_54R:HPe:2049:Z92",
		419,
		10,
		35,
	}
	PM_C_GP_9x54R_APFSS=cartridgeT{
		"9x54R",
		"High-output gauss cartridge firing kinetic sabots, meant to defeat medium body armor along with lower-end heavy kinetic protection, punches through most personnel shielding effectiely in bursts as if it were a superheated high-frequency tungsten knife slicing through butter.\nThe high velocities the sabot's fired at results in a thin coating of plasma to form around the penetrator, causing burns to exposed flesh and flammable materials.",
		9,
		54,
		66,
		7.62,
		37,
		64,
		4778,
		19,
		.0000127,
		1,
		"Armor-Piercing Fin-Stabilized Sabot - PM:C:GP:9_54R:APFSS:2047:Z201",
		532,
		11,
		12,
	}//subprojectileID 12 for the discarding cartridge, both 11 and 12 is fired unlike other cartridges with a listed subprojectileID >0
//Heavy (listed as 570x762 as i cant have decimals in variable names)
	PM_C_CQB_570x762R_Flechette16PD=cartridgeT{
		"57x76.2R",
		"High capacity flechette point-detonate cartridges - utilizes heavy, high-accuracy sabots.\nDeveloped for 57mm grenade launchers such as the automatic N799 line or the underbarrel/individual N255.",
		57,
		76.2,
		750,
		7.5,
		27,
		45.3,
		2200,
		76,
		.18,
		1,
		"WTi[YX]16 Heavy Sabot - Flechette Point Detonate - PM:C:CQB:57_76.2R:16:PD:2085:X27",
		119,
		12,
		13,
	}//12 is the main projectile, before fuse is setoff for subprojectile 13s
	PM_C_GP_570x762R_SMOKE=cartridgeT{
		"57x76.2R",
		"Smoke grenade capable of masking IR, visible light, and UV signatures.\nLasts 15 seconds.\nBaseline variant strictly for use with infantry.",
		57,
		76.2,
		223,
		56,
		66,
		200,
		1270,
		11,
		.05,
		1,
		"Smoke Grenade - PM:C:GP:57_76.2R:SMOKE:2035:Z355",
		25,
		14,
		15,
	}
	PM_C_GP_762_178__570_762R_AIRBURSTSMOKE=cartridgeT{
		"57x76.2R",
		"3in smoke grenade usually found on armored fighting vehicles and heavy combat frames -  Adapted for infantry use, thus the smoke grenade is external, with ferric propellant being downscaled to 57mm and contained within the launcher - the N799 cannot automatically load such grenades due to their design.\nDetonates a few meters above the ground with the use of a programmable EM fuze.",
		76.2,
		178,
		850,
		76.2,
		101.8,
		800,
		2100,
		13,
		.025,
		1,"Airburst Smoke Grenade - PM:C:GP:76.2:2058:E32",
		35,
		16,
		17,
	}//subprojectile 15 for smoke grenades?
	//insert EMP and HEDP im too lazy for this lol
	//90mm
	PM_C_AT_90_1270_2EFP_TN=cartridgeT{
		"90x1270",
		"Unguided, near-recoilless high-velocity anti-armor rocket-propelled grenade, utilizes a tandem low-yield thermonuclear Explosively Formed Penetrator warhead which essentially accelerates dense particles into a jet of high-velocity, dense superheated particles to punch through their target.",
		90,
		1270,
		16500,
		90,
		1200,
		16000,
		6500,
		16,
		.00000005,
		1,
		"90mm RPG - PM:C:AT:90_1270:2EFP:TN:2070:C32",
		0,
		18,
		19,
	}
	PM_C_AT_90_1270_2TN=cartridgeT{
		"90x1270",
		"Unguided, near-recoilless high-velocity anti-personnel rocket-propelled grenade, utilizes two extremely-low-yield thermonuclear warheads to disable particle/electromagnetic shielding and following up with a secondary detonation to ensure the target's eliminated.",
		90,
		1270,
		14000,
		90,
		1200,
		13000,
		6600,
		16,
		.00000005,
		1,
		"90mm RPG - PM:C:AT:90_1270:2TN:2070:C32",
		0,
		18,
		20,
	}
//low-end guidance 901778
	PM_C_AT_90_1778_2EFP_TN=cartridgeT{
		"90x1270",
		"Low-end guidance, near-recoilless high-velocity anti-armor rocket-propelled grenade, utilizes a tandem low-yield thermonuclear Explosively Formed Penetrator warhead which essentially accelerates dense particles into a jet of high-velocity, superheated particles to punch through their target.",
		90,
		1778,
		16500,
		90,
		1700,
		16000,
		7000,
		18,
		.00000005,
		1,
		"90mm RPG - PM:C:AT:90_1778:2EFP:TN:2071:C36",
		0,
		21,
		22,
	}
	PM_C_AT_90_1778_2TN=cartridgeT{
		"90x1270",
		"Low-end guidance, near-recoilless high-velocity anti-personnel rocket-propelled grenade, utilizes two extremely-low-yield thermonuclear warheads to disable particle/electromagnetic shielding and following up with a secondary detonation to ensure the target's eliminated.",
		90,
		1778,
		16500,
		90,
		1700,
		16000,
		7100,
		18,
		.00000005,
		1,
		"90mm RPG - PM:C:AT:90_1778:2TN:2071:C36",
		0,
		21,
		23,
	}
	//155x1524 Infantry-Portable ATGM
	PM_C_AT_155_1550_2EFP_TN=cartridgeT{
		"155x1550",
		"Top attack smart anti-tank munition, usually mounted on Infantry Fighting Vehicles - utilizes two low-yield thermonuclear Explosively Formed Penetrator warheads to accelerate dense particles to a high percentage of the speed of light to punch through heavy armor arrays, with a tertiary extremely-low yield thermonuclear warhead to ensure the desired target has been disabled.",
		155,
		1550,
		16000,
		155,
		1500,
		15900,
		7500,
		23,
		0,
		1,
		"155mm Tandem Anti-Tank Weapon - PM:C:AT:155_1524_:2EFP:TN:2060:N3",
		0,
		24,
		25,
	}
	//155x1550 Infantry-Portable EML
	PM_C_AT_155_1550KE=cartridgeT{
		"155x1550",
		"Kinetic penetrator launched through an infantry-portable electromagnetic launcher to punch through composite armors much more effectively comparted to 155mm or 90mm ATGMs.",
		155,
		1550,
		195100,
		57,
		1530,
		195000,
		12776,
		37,
		0,
		1,
		"WTi[YX] APFSDS - PM:C:AT:155_1550:K7389C7",
		15922700,
		26,
		27,
	}
	//if you couldnt tell i cant get enough of flechette ammunition
	PM_C_AT_155_1550FLECHETTEPD=cartridgeT{
		"155x1550",
		"High-Calibre kinetic weapon launched through an infantry-portable electromagnetic launcher - utilizes a airburst flechette warhead with numerous sub-sabots for anti-personnel effect.",
		155,
		1550,
		16500,
		155,
		1500,
		15750,
		12737,
		40,
		0,
		1,
		"WTi[YX] 155 CANISTER - PM:C:AT:155_1550:K1337D12",
		0,
		28,
		29,
	}
	//15x57
	PM_C_GP_15_57_HPe=cartridgeT{
		"15x57",
		"High-Calibre handcannon cartridge with a kinetic Hollow-Point-Equivalent warhead.",
		15,
		57,
		555,
		12.7,
		32,
		255,
		3500,
		37,
		1,
		1,
		"WTi[YX] HPe - PM:C:GP:15_57:2077:A2",
		1456,
		32,
		33,
	}
	PM_C_GP_15_57_APFSS=cartridgeT{
		"15x57",
		"High-Calibre handcannon cartridge utilizing a lethal kinetic long-rod penetrator.",
		15,
		57,
		755,
		12.7,
		37,
		450,
		4500,
		45,
		1,
		1,
		"WTi[YX] Armor-Piercing Fin-Stabilized Sabot - PM:C:GP:15_57:2130",
		3405,
		37,
		38,
	}
	//15x150
	PM_C_GP_15_150HPe=cartridgeT{
		"15x150",
		"Anti-material cartridge to punch through heavy kinetic armor and light armored vehicles, projectile equipped with Hollow-Point equivalent tip,",
		15,
		150,
		0,
		12.7,
		70,
		0,
		5323,
		57,
		1,
		1,
		"WTi[YX] Hollow-Point equivalent - PM:C:GP:15_150HP:2123A4",
		0,
		0,
		0,//PLACEHOLDER
	}
	//PM_C_GP_15_150APFSS
	//20x220 SUNBEAM
	PM_C_PD_20_220_APFSDS=cartridgeT{
		"20x220",
		"General-Purpose low-calibre autocannon cartridge, albeit usually utilized for extremely-short-ranged point-defense.\nFor the sake of backwards compatability with older, conventional propellant models, 20*220 cartridges have a 12.7mm diameter 5in sabot telescoped into the casing filled with ferromagnetic propellant triggered via plasma primer which in itself is triggered by an electric firing pin.",
		20,
		220,
		1550,
		15,
		127,
		1200,
		8355,
		90,
		0.0002,
		1,
		"WTi[YX] Armor-Piercing Fin-Stabilized Discarding-Sabot - PM:C:PD:20_220:2088:A5",
		41837.524,
		39,
		40,
	}
	///---Weaponry---///
	//wtfff not enough values apparently
	
	//N255 Individual Grenade Launcher
	PM_F_GP_570_762R_IND=weapon{
		255,true,
		"PM:F:N255:57_76.2:IND:2044:B20","57x76.2R","An individual grenade launcher with a short grip and rangefinder.",
		57,
		76.2,
		1270,
		240,
		.000000000000002,
		.95,
		11,
		13,
		255,255,255,255,255,
	}
	//N255 Attachment Variant
	PM_F_GP_570_762R=weapon{
		256, true,
		"PM:F:N255:57_76.2:2044:B20","57x76.2R","An individual grenade launcher attachment to be added onto the underbarrel hardpoint of other firearms, can be linked with said firearm's rangefinder if compatible.",
		57,
		76.2,
		800,
		240,
		.000000000000002,
		1,
		1,
		11,
		255,255,255,255,255,
	}//
	//N799 Automatic Grenade Launcher
	PM_F_GP_570_762R_AUTOGL=weapon{
		799,false,
		"PM:F:GP:N799:57_76.2R:2143:V9","57x76.2R","An automatic grenade launcher usually fitted on Armored Fighting Vehicles for urban combat, or heavy combat frames - standard variant usually equipped with a 57/9 barrel length, a digital rangefinder, night-vision compatible sights, a bipod, along with a physical rangefinder tucked away.\nUnable automatically fire 76.2x178 smoke grenades adapted for 57x76.2R grenade launchers - requires manual loading.",
		57,
		76.2,
		35500,
		360,
		.0000000000000023,
		1,
		1,
		12,
		799,799,799,799,799,
	}
	//N310 IN SERVICE SINCE 2027 RAHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH
	PM_F_GP_9_54R_BULLPUP=weapon{
		310,false,
		"PM:F:GP:N310-V:9_54R:2140:Y52","9x54R", "A highly ergonomic, bullpupped assault rifle for general use, excels at medium to intermediate-long-range combat - yet works quite well at most ranges.\nBaseline variant fitted with a 16in gauss accelerator with an integrated fuze-programmer at the muzzle in the case of conversion to 23x54R, along with night-vision compatible sights reminescent to some ancient scope known as the ACOG, includes a digital rangefinder, a bayonet carried within the stock, an angled grip, flippable composite sights, along with standard hardpoints for attachments.\nAmazing when fired in bursts of three to four rounds.",
		9,
		54,
		3500,
		975,
		.000000000000037,
		1,
		1,
		23,
		310,310,310,310,310,
	}
	PM_F_GP_9_54R_SMG=weapon{
		311,false,
		"PM:F:GP:N311-V4:9_54R:2133:Y32","9x54R","An ergonomic bullpupped SMG variant of the N310, excelling at short to medium range combat and quick infiltration operations.\nBase variant equipped with the N310's standard scope and sights along with the digital rangefinder, utilizes a shorter 8 inch barrel, along with its standard hardpoints.\nDue to the shorter barrel and thus increased recoil, it's recommended to only fire in bursts of three to four rounds.",
		9,
		54,
		2700,
		975,
		.0000000000000375,
		1.3,
		1,
		29,
		311,311,311,311,311,
	}
	PM_F_GP_9_54R_N310DASHN355_LMG=weapon{
		355,false,
		"PM:F:GP:N355-A2:9_54R_2128:Y23","9x54R","An LMG variant of the N310 rifle - the baseline variant equipped with the N310's standard scope and thus its digital rangefinder along with a bipod making the N355 great for suppressing hostiles and/or for continous firing.\nFor obvious reasons, it's recommended to not take this to close-quarters engagements.",
		9,
		54,
		5700,
		1050,
		.000000000000023,
		1,
		.9,
		33,
		355,355,355,355,355,
	}
	//10x25 cartridge
	PM_F_GP_10_25_N254=weapon{
		254,false,
		"PM:F:GP:N254:10_25:2140:A2","10x25","A low-calibre select-fire handgun, baseline firing a 10x25 gauss cartridge - relatively short effective range, yet trusty in the case you were to run out of ammunition.\nRemember - switching to your secondary is quicker than reloading.",
		10,
		25,
		610,
		720,
		.00000000000000254,
		1,
		.9,
		1,
		254,254,254,254,254,
	}
	PM_F_CQB_10_25_N312=weapon{
		312,false,
		"PM:F:CQB:N312-V5:10_25:2133:Y32","10x25","An adapatation of the N311 SMG - which in itself is a modification of the N310, which further downscales the SMG for a 10x25 pistol cartridge, albeit it retains the same accelerator length giving the cartridge more room to accelerate than a standard 10x25 pistole would grant.\nAmazing for CQB and quick infiltration operations, equipped with a digital rangefinder and the N310 family's standard scope to be able to perform decently even at medium ranges.",
		10,
		25,
		1890,
		975,
		.0000000000000378,
		1.3,
		1,
		1,
		29,312,312,312,312,//why is it 29ID
	}
	//15x57
	PM_F_GP_15_57_N555=weapon{
		555,false,
		"PM:F:GP:N555:15_57:2099:A7","15x57","A reworked version of the N155 handcannon which utilized non-gauss 15x45 cartridges, only recommended for individuals capable of handling the immense recoil as energy is pushed outwards from the firearm's intense accelerator and ferric propellant from the cartridge.",
		15,
		57,
		1200,
		660,
		0,
		1,
		1,
		1,
		555,555,555,555,555,
	}
	PM_F_CQB_15_57_N557=weapon{
		557,false,
		"PM:F:GP:N557:15_57:2122:A3","15x57","An overhaul of the N310 to accomodate the higher capacity 15x57 handcannon cartridge.\nFor obvious reasons, it's recommended you use equipment to significantly dampen the recoil even with semi-automatic fire if you are currently unable to.",
		15,
		57,
		4215,
		950,
		0,
		1,
		1,
		1,
		557,557,557,557,557,
	}
	//15x150
	PM_F_GP_15_150_N330=weapon{
		330,false,
		"PM:F:GP:N330:15_150:2089:Y2","15x150","A heavy nachine gun chambered in 15x150 - usually mounted on Armored Fighting Vehicles albeit this variant's equipped with a comparatively shorter barrel for easier portability with infantry.",
		15,
		150,
		7629,
		1270,
		.0000000000000434,
		1,
		1,
		1,
		330,330,330,330,330,
	}
	
	PM_F_GP_15_150_N333=weapon{
		333,false,
		"PM:F:GP:N333:15_150:2091:Y5","15x150","An anti-material rifle chambered in 15x150 - sufficiently long accelerator length allows this to punch through light armor with sabot ammunition, and more than enough for even the heaviest infantry body armor to be punched through even with particle/electromagnetic shielding in the way.",
		15,
		150,
		6800,
		480,
		0000000000000314,
		1,
		1,
		1,
		333,333,333,333,333,
	}
	
	//20x220 SUNBEAM Point-Defense Cannon
	PM_N_PD_20SUNBEAM_44=weapon{
		2044,false,
		"20/44 SUNBEAM PM:N:PD:20_220/44:2115:C3","20x220","\nA fourteen-barreled rotary autocannon chambered in the 20x220 SUNBEAM gauss cartridge - a 20/80 or 20/110 length is usually utilized on exoatmospheric naval vessels as the last line of defense against projectiles, or on aircraft for extremely-short-range combat - the 20/44 variant however is adapted for infantry use and meant for autonomous combat frames to handle - that, or users equipped with heavy power armor.\nEffective range dependant on recoil reduction, be advised as the 20/44 uses up a large amount of ammunition per burst.\n\n",
		20,
		220,
		215000,
		12754,
		.00000000000000014,
		1,
		1,
		1,
		2044,2044,2044,2044,2044,
	}
	
)

