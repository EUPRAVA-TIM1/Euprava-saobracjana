create database if not exists eupravaMilicija
CHARACTER SET utf8mb4
COLLATE utf8mb4_0900_ai_ci;
SET NAMES utf8mb4;

use eupravaMilicija;

    create table if not exists Opstina(
		PTT integer,
        Naziv varchar(20),
        primary key (ptt)
    );
    
    create table if not exists PolicijskaStanica(
		Id int auto_increment,
		Adresa varchar(60),
		BrojTelefona varchar(13),
		Email varchar(35),
        VremeOtvaranja varchar(8),
        VremeZatvaranja varchar(8),
		OpstinaID  integer,
		primary key(Id),
		foreign key (OpstinaId) references Opstina(PTT)
    );
    
    create table if not exists Zaposleni(
		JMBG varchar(13),
		RadiU int auto_increment,
		primary key(JMBG),
		foreign key (RadiU) references PolicijskaStanica(Id)
    );
    
    create table if not exists PrekrsajniNalog(
		Id int auto_increment,
        Datum date not null,
        Opis varchar(200) not null,
        IzdatoOdStrane varchar(40) not null,
        JMBGSluzbenika varchar(13) not null,
		IzdatoZa varchar(40) not null,
        JMBGZapisanog varchar(13) not null,
        TipPrekrsaja enum('POJAS','PREKORACENJE_BRZINE','PIJANA_VOZNJA','TEHNICKA_NEISPRAVNOST','PRVA_POMOC','NEMA_VOZACKU','REGISTRACIJA') not null,
        JedinicaMere enum ('promil','km/h'),
        Vrednost float,
        KaznaIzvrsena bool default false not null ,
		primary key (Id)
    );

    create table if not exists SlikeNaloga(
		NalogId int auto_increment,
        UrlSlike varchar(100) ,
        foreign key (NalogId) references PrekrsajniNalog(Id),
        primary key (NalogId,UrlSlike)
    );

    create table if not exists SudskiNalog(
		Id int auto_increment,
        Datum date not null,
        Naslov varchar(100) not null,
        Opis varchar(200) not null,
        IzdatoOdStrane varchar(40) not null,
        JMBGSluzbenika varchar(13) not null,
        Optuzeni varchar(40) not null,
        JMBGOptuzenog varchar(13) not null,
        StatusSlucaja enum ('POSLAT','U_PROCESU','ODBIJEN','PRESUDJEN','POTREBNI_DOKAZI') not null,
        primary key (Id)
    );
    
	create table if not exists DokumentiSudskogNaloga(
		NalogId int auto_increment,
        UrlDokumenta varchar(100),
        foreign key (NalogId) references SudskiNalog(Id),
        primary key (NalogId,UrlDokumenta)
    );

    create table if not exists Secrets(
        Id int auto_increment,
        ExpiresAt date,
        SecretKey varchar(64),
        primary key (Id)
    );

INSERT INTO Secrets (Id,ExpiresAt, SecretKey) VALUES (1,Now(),'secret');

-- Unos podataka o opštinama
INSERT INTO Opstina (PTT, Naziv) VALUES
(11000, 'Beograd'),
(21000, 'Novi Sad'),
(24000, 'Subotica'),
(18000, 'Niš'),
(36000, 'Kraljevo'),
(34000, 'Kragujevac');


-- Unos podataka o stanicama
INSERT INTO PolicijskaStanica (Id,Adresa,BrojTelefona,Email,VremeOtvaranja,VremeZatvaranja,OpstinaID) VALUES
(1, 'Knez Mihailova 10', '011/111-111', 'ps1@policija.rs', '08:00', '20:00', 11000),
(2, 'Bulevar Despota Stefana 15', '011/222-222', 'ps2@policija.rs', '07:00', '19:00', 21000),
(3, 'Kralja Petra 20', '011/333-333', 'ps3@policija.rs', '09:00', '21:00', 11000),
(4, 'Nemanjina 25', '011/444-444', 'ps4@policija.rs', '08:30', '19:30', 21000),
(5, 'Bulevar Kralja Aleksandra 30', '011/555-555', 'ps5@policija.rs', '10:00', '22:00', 21000),
(6, 'Vojvode Stepe 35', '011/666-666', 'ps6@policija.rs', '09:30', '20:30', 18000),
(7, 'Bulevar Mihajla Pupina 40', '011/777-777', 'ps7@policija.rs', '07:30', '18:30', 24000),
(8, 'Vuka Karadzica 45', '011/888-888', 'ps8@policija.rs', '08:00', '20:00', 36000),
(9, 'Takovska 50', '011/999-999', 'ps9@policija.rs', '10:30', '22:30', 11000),
(10, 'Resavska 55', '011/000-000', 'ps10@policija.rs', '07:00', '19:00', 34000);

-- Unos zaposlenih 
INSERT INTO Zaposleni (JMBG,RadiU) VALUES 
('0103974784007',1),
('1609943747006',1),
('2712969749004',3),
('1005998175053',2),
('0505967753001',2),
('2601952150049',4),
('0401961150046',5),
('1409930721015',6),
('2608959721013',6),
('0109984762012',7),
('2808986701006',7),
('2902939725016',8),
('0203988746002',8),
('1702970150021',9),
('2010960150022',10),
('0911963150049',10);

-- Unos prekrsajnih naloga
INSERT INTO PrekrsajniNalog ( Datum, Opis, IzdatoOdStrane, JMBGSluzbenika, IzdatoZa, JMBGZapisanog, TipPrekrsaja, JedinicaMere, Vrednost) VALUES
( '2023-05-04', 'Vožnja bez pojasa', 'Mila Milic', '1609943747006', 'Katarina Katarinic', '0101993175038', 'POJAS', null, null),
( '2023-05-03', 'Prekoračenje brzine za 20km/h', 'Mila Milic', '1609943747006', 'Katarina Katarinic', '0101993175038', 'PREKORACENJE_BRZINE', 'km/h', 20),
( '2023-04-30', 'Pijana vožnja', 'Mila Milic', '1609943747006', 'Katarina Katarinic', '0101993175038', 'PIJANA_VOZNJA', 'promil', 1),
( '2023-04-29', 'Tehnička neispravnost vozila', 'Petar Petkovic', '0911963150049', 'Ivan Ivanovic', '1507993175075', 'TEHNICKA_NEISPRAVNOST', null, null),
( '2023-04-28', 'Nedavanje prve pomoći', 'Petar Petkovic', '0911963150049', 'Ivan Ivanovic', '1507993175075', 'PRVA_POMOC', null, null),
( '2022-02-14', 'Vožnja sa neispravnom kočnicom', 'Petar Petkovic', '0911963150049', 'Ivan Ivanovic', '1507993175075', 'TEHNICKA_NEISPRAVNOST', null, null),
( '2022-04-20', 'Nevezivanje pojasa tokom vožnje', 'Goran Goranic', '1702970150021', 'Hana Hanic', '2012995175033', 'POJAS', null, null),
( '2022-05-01', 'Vožnja sa isteklom registracijom', 'Goran Goranic', '1702970150021', 'Hana Hanic', '2012995175033', 'REGISTRACIJA', null, null),
( '2022-06-10', 'Upotreba mobilnog telefona tokom vožnje', 'Goran Goranic', '1702970150021', 'Hana Hanic', '2012995175033', 'TEHNICKA_NEISPRAVNOST', null, null),
( '2022-07-15', 'Prekoračenje brzine za 50 km/h', 'Goran Goranic', '1702970150021', 'Nikola Nikolić', '2004993175065', 'PREKORACENJE_BRZINE', 'km/h', 50),
( '2022-02-14', 'Vožnja sa neispravnom kočnicom', 'Ana Anic', '2712969749004', 'Eva Evic', '2004993175065', 'TEHNICKA_NEISPRAVNOST', null, null),
( '2022-04-20', 'Nevezivanje pojasa tokom vožnje', 'Ana Anić', '2712969749004', 'Eva Evic', '2004993175065', 'POJAS', null, null),
( '2022-05-01', 'Vožnja sa isteklom registracijom', 'Ana  Anic', '2712969749004', 'Katarina Katarinic', '0101993175038', 'REGISTRACIJA', null, null),
( '2022-06-10', 'Upotreba mobilnog telefona tokom vožnje', 'Hana Hanic', '0505967753001', 'Katarina Katarinic', '0101993175038', 'TEHNICKA_NEISPRAVNOST', null, null),
( '2022-07-15', 'Prekoračenje brzine za 50 km/h', 'Hana Hanic', '0505967753001', 'Katarina Katarinic', '0101993175038', 'PREKORACENJE_BRZINE', 'km/h', 50);







