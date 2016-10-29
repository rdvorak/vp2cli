package vp2cli


type PacketParser {
	field string,
	offset int,
	length int,
	value func(b []byte) interface{}

"DateStamp" ,0 ,2
"TimeStamp" ,2 ,2
"OutsideTemperature" ,4 ,2  
"HighOutTemperature" ,6 ,2  
"LowOutTemperature" ,8 ,2
"Rainfall" ,10 ,2
"HighRainRate" ,12 ,2
"Barometer" ,14 ,2      
"SolarRadiation" ,16 ,2         
"NumberofWindSamples" ,18 ,2        
"InsideTemperature" ,20 ,2      
"InsideHumidity" ,22 ,1         
"OutsideHumidity" ,23 ,1        
"AverageWindSpeed" ,24 ,1
"HighWindSpeed" ,25 ,1
"DirectionofHiWindSpeed" ,26 ,1
"PrevailingWindDirection" ,27 ,1
"AverageUVIndex" ,28 ,1
"ET" ,29 ,1
"HighSolarRadiation" ,30 ,2
"HighUVIndex" ,32 ,1
"ForecastRule" ,33 ,1
"LeafTemperature" ,34 ,2
"LeafWetnesses" ,36 ,2
"SoilTemperatures" ,38 ,4
"DownloadRecordType" ,42 ,1
"ExtraHumidities" ,43 ,2
"ExtraTemperatures" ,45 ,3

